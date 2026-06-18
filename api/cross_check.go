package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
)

const apiURL = "https://qubexsentinel-api.onrender.com/verify"

type verifyReq struct {
	PublicKey string `json:"public_key"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type verifyResp struct {
	Valid bool `json:"valid"`
}

func callAPI(pk, msg, sig []byte) (bool, error) {
	body, _ := json.Marshal(verifyReq{
		PublicKey: hex.EncodeToString(pk),
		Message:   hex.EncodeToString(msg),
		Signature: hex.EncodeToString(sig),
	})
	resp, err := http.Post(apiURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	var v verifyResp
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return false, err
	}
	return v.Valid, nil
}

func runCase(name string, pk *mldsa87.PublicKey, msg, sig []byte, expected bool) {
	pkBytes, _ := pk.MarshalBinary()
	localResult := mldsa87.Verify(pk, msg, nil, sig)
	apiResult, err := callAPI(pkBytes, msg, sig)
	status := "PASS"
	if err != nil {
		status = "ERROR: " + err.Error()
	} else if localResult != apiResult || localResult != expected {
		status = "MISMATCH"
	}
	fmt.Printf("[%s] expected=%v circl_direct=%v api=%v -> %s\n",
		name, expected, localResult, apiResult, status)
}

func main() {
	pk, sk, _ := mldsa87.GenerateKey(rand.Reader)
	msg := []byte("cross-check message one")
	sig, _ := sk.Sign(rand.Reader, msg, nil)
	runCase("genuine signature", pk, msg, sig, true)

	tamperedSig := append([]byte{}, sig...)
	tamperedSig[0] ^= 0xFF
	runCase("tampered signature", pk, msg, tamperedSig, false)

	wrongMsg := []byte("a different message entirely")
	runCase("wrong message", pk, wrongMsg, sig, false)

	pk2, sk2, _ := mldsa87.GenerateKey(rand.Reader)
	sig2, _ := sk2.Sign(rand.Reader, msg, nil)
	runCase("signature from wrong key", pk, msg, sig2, false)
	_ = pk2
}
