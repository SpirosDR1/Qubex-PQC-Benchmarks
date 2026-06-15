package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
)

type verifyRequest struct {
	PublicKey string `json:"public_key"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type verifyResponse struct {
	Valid       bool   `json:"valid"`
	Scheme      string `json:"scheme"`
	VerifiedAt  string `json:"verified_at"`
	Statement   string `json:"statement,omitempty"`
	Attestation string `json:"attestation,omitempty"`
	Error       string `json:"error,omitempty"`
}

type attestRequest struct {
	Statement   string `json:"statement"`
	Attestation string `json:"attestation"`
}

type attestResponse struct {
	Genuine bool   `json:"genuine"`
	Scheme  string `json:"scheme"`
	Error   string `json:"error,omitempty"`
}

type rootResponse struct {
	Service   string   `json:"service"`
	Scheme    string   `json:"scheme"`
	Endpoints []string `json:"endpoints"`
}

var attestPK *mldsa87.PublicKey
var attestSK *mldsa87.PrivateKey

func main() {
	pk, sk, err := mldsa87.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal("attestation key generation failed: ", err)
	}
	attestPK, attestSK = pk, sk

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/verify", verifyHandler)
	http.HandleFunc("/verify-attestation", verifyAttestationHandler)
	http.HandleFunc("/attestation-key", attestationKeyHandler)
	http.HandleFunc("/demo", demoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Qubex Sentinel verification API listening on %s", addr)
	log.Printf("Scheme: ML-DSA-87 (NIST FIPS 204, level 5)")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, rootResponse{
		Service:   "Qubex Sentinel Verification API",
		Scheme:    "ML-DSA-87 (NIST FIPS 204, level 5)",
		Endpoints: []string{"/health", "/verify (POST)", "/verify-attestation (POST)", "/attestation-key", "/demo"},
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"scheme": "ML-DSA-87",
	})
}

func attestationKeyHandler(w http.ResponseWriter, r *http.Request) {
	pkBytes, err := attestPK.MarshalBinary()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "marshal failed"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"scheme":                 "ML-DSA-87",
		"attestation_public_key": hex.EncodeToString(pkBytes),
		"statement_format":       "qubex-attestation-v1|scheme=ML-DSA-87|pk_sha256=<hex>|msg_sha256=<hex>|result=<valid|invalid>|at=<RFC3339>",
		"note":                   "Attestation key is currently ephemeral and rotates on restart. Persistence is on the roadmap.",
	})
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, verifyResponse{Error: "use POST"})
		return
	}

	var req verifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, verifyResponse{Error: "invalid JSON body"})
		return
	}

	pkBytes, err := hex.DecodeString(req.PublicKey)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, verifyResponse{Error: "public_key is not valid hex"})
		return
	}
	msgBytes, err := hex.DecodeString(req.Message)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, verifyResponse{Error: "message is not valid hex"})
		return
	}
	sigBytes, err := hex.DecodeString(req.Signature)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, verifyResponse{Error: "signature is not valid hex"})
		return
	}

	var pk mldsa87.PublicKey
	if err := pk.UnmarshalBinary(pkBytes); err != nil {
		writeJSON(w, http.StatusBadRequest, verifyResponse{Error: "public_key has wrong length for ML-DSA-87"})
		return
	}

	valid := mldsa87.Verify(&pk, msgBytes, nil, sigBytes)

	at := time.Now().UTC().Format(time.RFC3339)
	resp := verifyResponse{
		Valid:      valid,
		Scheme:     "ML-DSA-87",
		VerifiedAt: at,
	}

	pkHash := sha256.Sum256(pkBytes)
	msgHash := sha256.Sum256(msgBytes)
	result := "invalid"
	if valid {
		result = "valid"
	}
	statement := "qubex-attestation-v1|scheme=ML-DSA-87|pk_sha256=" +
		hex.EncodeToString(pkHash[:]) + "|msg_sha256=" +
		hex.EncodeToString(msgHash[:]) + "|result=" + result + "|at=" + at

	attSig, signErr := attestSK.Sign(rand.Reader, []byte(statement), nil)
	if signErr == nil {
		resp.Statement = statement
		resp.Attestation = hex.EncodeToString(attSig)
	}

	writeJSON(w, http.StatusOK, resp)
}

func verifyAttestationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, attestResponse{Error: "use POST"})
		return
	}

	var req attestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, attestResponse{Error: "invalid JSON body"})
		return
	}

	if req.Statement == "" || req.Attestation == "" {
		writeJSON(w, http.StatusBadRequest, attestResponse{Error: "statement and attestation are required"})
		return
	}

	attBytes, err := hex.DecodeString(req.Attestation)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, attestResponse{Error: "attestation is not valid hex"})
		return
	}

	genuine := mldsa87.Verify(attestPK, []byte(req.Statement), nil, attBytes)

	writeJSON(w, http.StatusOK, attestResponse{
		Genuine: genuine,
		Scheme:  "ML-DSA-87",
	})
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	pk, sk, err := mldsa87.GenerateKey(rand.Reader)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "keygen failed"})
		return
	}

	msg := []byte("Qubex Sentinel verification demo")
	sig, err := sk.Sign(rand.Reader, msg, nil)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "signing failed"})
		return
	}

	pkBytes, err := pk.MarshalBinary()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "marshal failed"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"public_key": hex.EncodeToString(pkBytes),
		"message":    hex.EncodeToString(msg),
		"signature":  hex.EncodeToString(sig),
	})
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
