package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
)

type vector struct {
	count int
	pk    string
	msg   string
	mlen  int
	sm    string
	smlen int
	ctx   string
}

func parseFile(path string) ([]vector, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var vectors []vector
	cur := vector{count: -1}

	flush := func() {
		if cur.count >= 0 && cur.pk != "" && cur.sm != "" {
			vectors = append(vectors, cur)
		}
	}

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024) // lines can be very long (multi-KB hex)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " = ", 2)
		if len(parts) != 2 {
			continue
		}
		key, val := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

		if key == "count" {
			flush()
			n, _ := strconv.Atoi(val)
			cur = vector{count: n}
			continue
		}

		switch key {
		case "pk":
			cur.pk = val
		case "msg":
			cur.msg = val
		case "mlen":
			cur.mlen, _ = strconv.Atoi(val)
		case "sm":
			cur.sm = val
		case "smlen":
			cur.smlen, _ = strconv.Atoi(val)
		case "ctx":
			cur.ctx = val
		}
	}
	flush()

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return vectors, nil
}

func main() {
	path := flag.String("file", "", "path to the .rsp KAT file")
	maxVectors := flag.Int("max", 0, "limit number of vectors tested (0 = all)")
	flag.Parse()

	if *path == "" {
		fmt.Println("usage: go run kat_check.go -file kat_MLDSA_87_det_pure.rsp")
		return
	}

	vectors, err := parseFile(*path)
	if err != nil {
		fmt.Println("failed to read/parse file:", err)
		return
	}
	if len(vectors) == 0 {
		fmt.Println("no vectors parsed — check the file path and format")
		return
	}

	fmt.Printf("Parsed %d test vectors from %s\n\n", len(vectors), *path)

	total := len(vectors)
	if *maxVectors > 0 && *maxVectors < total {
		total = *maxVectors
	}

	pass, fail := 0, 0

	for i := 0; i < total; i++ {
		v := vectors[i]

		pkBytes, err1 := hex.DecodeString(v.pk)
		smBytes, err2 := hex.DecodeString(v.sm)
		ctxBytes, err3 := hex.DecodeString(v.ctx)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Printf("[count=%d] hex decode error, skipping\n", v.count)
			fail++
			continue
		}

		if len(smBytes) != v.smlen || v.smlen < v.mlen {
			fmt.Printf("[count=%d] length mismatch (smlen=%d, got %d bytes), skipping\n", v.count, v.smlen, len(smBytes))
			fail++
			continue
		}

		sigLen := v.smlen - v.mlen
		sigBytes := smBytes[:sigLen]
		msgBytes := smBytes[sigLen:]

		var pk mldsa87.PublicKey
		if err := pk.UnmarshalBinary(pkBytes); err != nil {
			fmt.Printf("[count=%d] pk unmarshal error: %v\n", v.count, err)
			fail++
			continue
		}

		valid := mldsa87.Verify(&pk, msgBytes, ctxBytes, sigBytes)
		if valid {
			pass++
		} else {
			fail++
			fmt.Printf("[count=%d] FAIL — expected valid signature, got invalid\n", v.count)
		}
	}

	fmt.Println()
	fmt.Println("=== Results ===")
	fmt.Printf("Tested:  %d\n", total)
	fmt.Printf("Passed:  %d\n", pass)
	fmt.Printf("Failed:  %d\n", fail)
	if fail == 0 {
		fmt.Println("\nAll official NIST-aligned ML-DSA-87 test vectors verified correctly.")
	}
}
