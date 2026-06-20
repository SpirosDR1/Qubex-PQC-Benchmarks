package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	baseURL    = "https://qubexsentinel-api.onrender.com"
	demoPath   = "/demo"
	verifyPath = "/verify"
)

type demoPayload struct {
	PublicKey string `json:"public_key"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

func fetchDemoPayload() (demoPayload, error) {
	var d demoPayload
	resp, err := http.Get(baseURL + demoPath)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(body, &d)
	return d, err
}

func sendVerify(payload []byte) (time.Duration, bool, error) {
	start := time.Now()
	resp, err := http.Post(baseURL+verifyPath, "application/json", bytes.NewReader(payload))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, false, err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)
	return elapsed, resp.StatusCode == 200, nil
}

func percentile(sorted []time.Duration, p float64) time.Duration {
	if len(sorted) == 0 {
		return 0
	}
	idx := int(p * float64(len(sorted)-1))
	return sorted[idx]
}

func main() {
	concurrency := flag.Int("concurrency", 10, "number of concurrent workers")
	totalRequests := flag.Int("requests", 100, "total number of requests to send")
	flag.Parse()

	fmt.Println("Warming up: fetching one demo payload to reuse across all requests...")
	demo, err := fetchDemoPayload()
	if err != nil {
		fmt.Println("Failed to fetch demo payload (API may be cold-starting, try again in ~60s):", err)
		return
	}
	payload, _ := json.Marshal(demo)
	fmt.Printf("Got demo payload. Running load test: %d requests, %d concurrent workers...\n\n", *totalRequests, *concurrency)

	var (
		mu         sync.Mutex
		latencies  []time.Duration
		successes  int
		failures   int
		errorsList []string
	)

	sem := make(chan struct{}, *concurrency)
	var wg sync.WaitGroup

	overallStart := time.Now()

	for i := 0; i < *totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			elapsed, ok, err := sendVerify(payload)

			mu.Lock()
			defer mu.Unlock()
			latencies = append(latencies, elapsed)
			if err != nil {
				failures++
				errorsList = append(errorsList, err.Error())
			} else if ok {
				successes++
			} else {
				failures++
			}
		}()
	}

	wg.Wait()
	overallElapsed := time.Since(overallStart)

	sort.Slice(latencies, func(i, j int) bool { return latencies[i] < latencies[j] })

	var sum time.Duration
	for _, l := range latencies {
		sum += l
	}
	mean := time.Duration(0)
	if len(latencies) > 0 {
		mean = sum / time.Duration(len(latencies))
	}

	fmt.Println("=== Results ===")
	fmt.Printf("Total requests:     %d\n", *totalRequests)
	fmt.Printf("Successful:         %d\n", successes)
	fmt.Printf("Failed:             %d\n", failures)
	fmt.Printf("Total wall time:    %v\n", overallElapsed)
	if overallElapsed.Seconds() > 0 {
		fmt.Printf("Throughput:         %.2f requests/sec\n", float64(*totalRequests)/overallElapsed.Seconds())
	}
	fmt.Println()
	fmt.Printf("Latency mean:       %v\n", mean)
	if len(latencies) > 0 {
		fmt.Printf("Latency min:        %v\n", latencies[0])
		fmt.Printf("Latency p50:        %v\n", percentile(latencies, 0.50))
		fmt.Printf("Latency p95:        %v\n", percentile(latencies, 0.95))
		fmt.Printf("Latency p99:        %v\n", percentile(latencies, 0.99))
		fmt.Printf("Latency max:        %v\n", latencies[len(latencies)-1])
	}

	if failures > 0 && len(errorsList) > 0 {
		fmt.Println("\nSample errors:")
		max := 5
		if len(errorsList) < max {
			max = len(errorsList)
		}
		for _, e := range errorsList[:max] {
			fmt.Println(" -", e)
		}
	}
}
