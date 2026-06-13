package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
)

func main() {
	const iterations = 100000

	fmt.Println("======================================================")
	fmt.Println("QUBEX SENTINEL | ML-DSA-87 PERFORMANCE BENCHMARK")
	fmt.Println("Scheme: NIST FIPS 204 (ML-DSA-87, Security Level 5)")
	fmt.Println("Library: cloudflare/circl")
	fmt.Printf("Iterations per op: %d\n", iterations)
	fmt.Println("======================================================")

	pk, sk, err := mldsa87.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Key generation failed:", err)
		return
	}

	msg := []byte("QUBEX_BENCHMARK_PAYLOAD")

	for i := 0; i < 1000; i++ { // warm-up
		sk.Sign(nil, msg, nil)
	}

	startSign := time.Now()
	for i := 0; i < iterations; i++ {
		sk.Sign(nil, msg, nil)
	}
	avgSign := time.Since(startSign).Nanoseconds() / int64(iterations)

	sig, err := sk.Sign(nil, msg, nil)
	if err != nil {
		fmt.Println("Signing failed:", err)
		return
	}

	startVerify := time.Now()
	for i := 0; i < iterations; i++ {
		mldsa87.Verify(pk, msg, nil, sig)
	}
	avgVerify := time.Since(startVerify).Nanoseconds() / int64(iterations)

	valid := mldsa87.Verify(pk, msg, nil, sig)

	fmt.Printf("Avg sign latency:   %d ns\n", avgSign)
	fmt.Printf("Avg verify latency: %d ns\n", avgVerify)
	fmt.Printf("Signature valid:    %v\n", valid)
	fmt.Println("======================================================")
}
