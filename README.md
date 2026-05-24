# QUBEX SENTINEL | PQC Performance Benchmark

A high-performance audit tool designed to measure the latency impact of Post-Quantum Cryptography (PQC) on EVM-compatible networks. This utility implements ML-DSA-44 (FIPS 204) signing and utilizes on-chain verification to ensure transparency in performance metrics.

### Why QUBEX SENTINEL?
As blockchain infrastructure transitions toward quantum resistance, understanding the cryptographic overhead is critical. This tool provides a deterministic way to benchmark signature latency, ensuring that PQC integration doesn't bottleneck high-frequency sequencer performance.

### Technical Specs
* Algorithm: ML-DSA-44 (FIPS 204 standard).
* Environment: Live integration with Base Sepolia (Testnet).
* Metrics:
    * Mean signing latency (in nanoseconds).
    * On-chain state verification (via block header synchronization).
    * Persistent audit logging for reproducibility.

### QUBEX SENTINEL | Multi-Chain Performance Matrix (NIST ML-DSA)

| Network | Type | Avg Signing Latency (ns) | Verification | Status |
| :--- | :--- | :--- | :--- | :--- |
| Mantle | Optimistic Rollup | 39,032 ns | On-chain | Ready |
| Metis | Optimistic Rollup | 39,499 ns | On-chain | Ready |
| Polygon (Amoy) | zkEVM / AggLayer | 39,616 ns | On-chain | Ready |
| Optimism | Optimistic Rollup | 39,712 ns | On-chain | Ready |
| Base | Optimistic Rollup | 68,244 ns | On-chain | Ready |
| BNB Smart Chain | L1 / Sidechain | 68,630 ns | On-chain | Ready |
| Linea | zkEVM | 93,278 ns | On-chain | Ready |
| Arbitrum | Optimistic Rollup | 122,314 ns | On-chain | Ready |
| Scroll | zkEVM | 122,538 ns | On-chain | Ready |
| zkSync | ZK-Rollup | 149,699 ns | On-chain | Ready |
| Blast | Optimistic Rollup | 230,952 ns | On-chain | Ready |

*Note: Benchmarks represent live ML-DSA (Dilithium) cryptographic stress tests directly interacting with respective network RPCs.*

### Quick Start

Prerequisites:
Ensure you have Go 1.26.3 or higher installed on your machine.

### Install dependencies:

```go
git clone [https://github.com/SpirosDR1/Qubex-PQC-Benchmark.git]
cd Qubex-PQC-Benchmark
go mod tidy
```

### Run the Stress Test:

```go
# Optional: Set a custom RPC URL for better performance
# export RPC_URL=https://your-alchemy-node-link

go run main.go [network_name]
```
*Supported networks: base, polygon, arbitrum, optimism, bnb, mantle, blast, zksync, linea, metis, scroll* 

### Expected Output & Audit Trail
The tool outputs real-time signing latency and verifies connectivity to the specified network.

### Terminal Output

```text
--- QUBEX SENTINEL | MULTI-CHAIN PQC BENCHMARK ---
[SYSTEM] Targeting Network: polygon
[SYSTEM] Starting PQC Stress Test: 100000 iterations...
--------------------------------------------------
[AUDIT] Infrastructure: QUBEX SENTINEL ENGINE
[AUDIT] Target Network: polygon
[AUDIT] Avg PQC Signing Latency: 39616 ns
[AUDIT] Signature Integrity: true (NIST ML-DSA Verified)
[AUDIT] Verified Block Height: 38887690
[AUDIT] Status: ALL CHECKS PASSED
--------------------------------------------------
QUBEX SENTINEL: polygon Ecosystem is Ready for PQC Deployment.
```

### Secure Audit Logging:

Each run automatically appends verifiable data to a network-specific log file (e.g., qubex_polygon_audit.log or qubex_base_audit.log). The audit trail provides an immutable-like record including:

* Execution Timestamp (RFC3339)
* Targeted Network
* Average PQC Latency (ns)
* Verified Blockchain Height
* Signature Verification Status

### Developed for QUBEX SENTINEL Research Initiative
License: MIT

