package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
)

func main() {
	pk, sk, err := mldsa87.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	pkBytes, err := pk.MarshalBinary()
	if err != nil {
		panic(err)
	}
	skBytes, err := sk.MarshalBinary()
	if err != nil {
		panic(err)
	}
	fmt.Println("ATTEST_PK_B64=" + base64.StdEncoding.EncodeToString(pkBytes))
	fmt.Println("ATTEST_SK_B64=" + base64.StdEncoding.EncodeToString(skBytes))
}
