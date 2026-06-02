# QUBEX SENTINEL: AI-Driven Post-Quantum DePIN Shield

The First Decentralized AI-PQC Middleware for EVM Ecosystems & Cross-Chain Bridges

## The PQC Omnichain Shield

The industry's first concurrent Post-Quantum Cryptography (PQC) middleware, securing EVM ecosystems against "Harvest Now, Decrypt Later" (HNDL) vectors.

## The "Quantum Collapse" & HNDL Threat
Legacy ECDSA rollups are exposed to an existential threat. State-level actors are currently executing "Harvest Now, Decrypt Later" (HNDL) attacks—archiving encrypted L1/L2 state data to decrypt it retroactively once Cryptographically Relevant Quantum Computers (CRQCs) reach maturity. The choice for modern ecosystems is binary: natively proof your infrastructure today, or accept the inevitability of total TVL extraction.
 
## Core Architecture: The "Zero L1 Impact" Paradigm
QUBEX Sentinel functions as a non-invasive, high-throughput interceptor layer positioned between the Sequencer and the Batcher. We secure networks without modifying existing consensus logic, avoiding the unsustainable gas inflation and state bloat caused by on-chain PQC implementations.

### Our Decoupled Pre-Batcher Execution Flow:

AI Threat Detection (Chaos Engine): Real-time interception of mempool traffic. Our predictive AI engine identifies and mitigates anomalous sequencing patterns indicative of HNDL vectors.

Decentralized Sentinel Nodes: Traffic is validated via our decentralized node network, utilizing NIST-standard ML-DSA-87 (Level 5) signatures with sub-millisecond mathematical validation.

Zero-Overhead Compression: Only cryptographically verified, quantum-resistant payloads reach the operator.

Result: Absolute PQC integrity on L2 with zero additional bytes of overhead on L1 calldata. We match legacy ECDSA throughput speeds, post-quantum.

## Live Omnichain Benchmark (Mainnet-Ready Deployment)

Validated via the QUBEX Omnichain Broadcaster on enterprise-grade infrastructure.

Security Standard: NIST Level 5 (ML-DSA-87).

Target Infrastructure: 8 EVM Ecosystems -> 6 Tier-1 L2 Rollups (Arbitrum, Optimism, Base, Scroll, Mantle, Polygon) + 2 L1 Settlement Layers (Ethereum, BNB Chain)

Validation Latency: Sub-millisecond (avg. ~170µs - 650µs).

Efficiency: 0% L1 gas bloat; 2.2 Million concurrent operations verified.

### Visual Proof

<img width="1152" height="815" alt="image" src="https://github.com/user-attachments/assets/23aaf3db-b918-45b0-aca3-ba30239f187b" />

<img width="1132" height="487" alt="image" src="https://github.com/user-attachments/assets/f45d4e84-bdb2-4059-9842-57c665273aa1" />


### Omnichain Reproduction & Local Verification
"Don't trust, verify." QUBEX Sentinel operates with absolute cryptographic transparency. Enterprise architects and protocol core developers can independently reproduce the stress test locally.

#### Prerequisites
- Go Environment: Version 1.20 or higher.
- A burner EVM wallet with testnet gas across the 8 networks.
- Hardware: Multi-core CPU. *(Warning: The benchmark is designed to force 100% CPU thread exhaustion to simulate peak enterprise sequencer loads. Proceed with caution on commercial laptops)*.

#### Execution Protocol

```go
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
```

```go
cd Qubex-PQC-Benchmarks/omnichain-router
```

```go
go mod tidy
```

### Inject your burner wallet's private key (NO 0x prefix)

#### Linux/Mac:

export PRIVATE_KEY="YOUR_BURNER_PRIVATE_KEY"

#### Windows (PowerShell):

$env:PRIVATE_KEY="YOUR_BURNER_PRIVATE_KEY"

#### Ignite the Devnet Broadcaster

```go
go run tx.go
```

## Integration & Genesis Pilots

QUBEX Sentinel is currently in the MVP Phase, actively selecting high-TVL Cross-Chain Bridges, Oracles, and Tier-1 L2 Rollups for exclusive Genesis Pilot integrations.

Networks participating in the Genesis Pilot phase are eligible for Strategic Integration Bounties (subsidized from our Ecosystem Grants allocation) to eliminate the financial friction of transitioning to a quantum-safe architecture.

### Initiate Integration:

To request a technical handoff, review the Trust Index framework, and secure a Genesis Pilot slot, submit your protocol's architecture for review by our core team.

Institutional Contact: spyridongagr@qubexsentinel.com

X (Twitter): @QUBEX_SENTINEL

[Submit Genesis Pilot Request Here](https://forms.gle/hmUdBiQz3PT2x8TT7)
