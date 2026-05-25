# QUBEX SENTINEL | PQC Performance Benchmarks (Chaos Engine)

A high-performance audit tool designed to measure the latency impact of Post-Quantum Cryptography (PQC) on EVM-compatible networks. This utility implements NIST ML-DSA-44 (FIPS 204) signing and utilizes on-chain verification to ensure transparency in performance metrics.

### Live Portals & Genesis Program

* Official Portal: [qubexsentinel.com](https://qubexsentinel.com)

* Live Execution Logs: [qubexsentinel.com/benchmarks](https://qubexsentinel.com/benchmarks)

* Apply for Genesis Partnership (Zero-Fee Integration): [Access the Application Form Here](https://forms.gle/hmUdBiQz3PT2x8TT7)

---

## Why QUBEX SENTINEL?

As blockchain infrastructure transitions toward quantum resistance to combat HNDL (Harvest Now, Decrypt Later) attacks, understanding cryptographic overhead is critical. The Chaos Engine provides a deterministic way to benchmark signature latency, ensuring that PQC integration introduces virtually zero drag to high-frequency sequencer performance.

## Multi-Chain Performance Matrix (NIST ML-DSA)

*Verified metrics interacting directly with respective network RPCs (May 2026).*

| Rank | Network         | Type              | Avg Signing Latency | Verification | Status |
| ---- | --------------- | ----------------- | ------------------- | ------------ | ------ |
| 1    | Base            | Optimistic Rollup |           39,595 ns | On-chain     | Ready  |
| 2    | Arbitrum        | Optimistic Rollup |           65,536 ns | On-chain     | Ready  |
| 3    | Scroll          | zkEVM             |           91,899 ns | On-chain     | Ready  |
| 4    | Linea           | zkEVM             |          118,647 ns | On-chain     | Ready  |
| 5    | BNB Smart Chain | L1 / Sidechain    |          121,210 ns | On-chain     | Ready  |
| 6    | Mantle          | Optimistic Rollup |          124,309 ns | On-chain     | Ready  |
| 7    | zkSync          | ZK Rollup         |          127,895 ns | On-chain     | Ready  |
| 8    | Blast           | Optimistic Rollup |          181,195 ns | On-chain     | Ready  |
| 9    | Metis           | Optimistic Rollup |          187,759 ns | On-chain     | Ready  |
| 10   | Optimism        | Optimistic Rollup |          246,632 ns | On-chain     | Ready  |
| 11   | Polygon (Amoy)  | zkEVM / AggLayer  |          438,600 ns | On-chain     | Ready  |


## Quick Start (Verify it yourself)

Prerequisites: Ensure you have Go installed on your machine.

1. Clone the repository:
```go
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
cd Qubex-PQC-Benchmarks
go mod tidy
```
2. Run the Chaos Engine Stress Test:

```go
go run main.go polygon
``` 

*Supported networks: base, polygon, arbitrum, optimism, bnb, mantle, blast, zksync, linea, metis, scroll*

## Expected Terminal Output & Audit Trail

The tool outputs real-time signing latency and verifies connectivity to the specified network.

```go
--- QUBEX SENTINEL | MULTI-CHAIN PQC BENCHMARK ---

[SYSTEM] Targeting Network: polygon

[SYSTEM] Starting PQC Stress Test: 100000 iterations...

[AUDIT] Infrastructure: QUBEX SENTINEL ENGINE

[AUDIT] Target Network: polygon

[AUDIT] Avg PQC Signing Latency: 39616 ns

[AUDIT] Signature Integrity: true (NIST ML-DSA Verified)

[AUDIT] Verified Block Height: 38887690

[AUDIT] Status: ALL CHECKS PASSED

--------------------------------------------------

QUBEX SENTINEL: polygon Ecosystem is Ready for PQC Deployment.
```

## Secure Audit Logging

Each run automatically appends verifiable data to a network-specific log file (e.g., qubex_polygon_audit.log). The audit trail provides an immutable-like record including:

* Execution Timestamp (RFC3339)

* Targeted Network

* Average PQC Latency (ns)

* Verified Blockchain Height

* Signature Verification Status

## Architecture & IP Note

*Pro Tip: For high-throughput L2/L3 testing, ensure your environment supports high-concurrency Go execution as decoupled by our production engine.*

*Security Note: This public repository contains our high-performance stress-testing engine (Go). The core PQC cryptographic logic (Python/NIST Level 5 implementations) is maintained in a private repository to protect proprietary IP. Full white-box access is granted only under NDA during commercial integration.*

---

## Ready for Production?

We are actively selecting Tier-1 RaaS providers and L2/L3 ecosystems for the Genesis Partnership Program (6-month zero-fee integration).

[Initiate Integration via the Genesis Form](https://forms.gle/hmUdBiQz3PT2x8TT7)

*Developed for QUBEX SENTINEL Research Initiative* 

*License: MIT*
