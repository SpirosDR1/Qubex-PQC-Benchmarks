# QUBEX SENTINEL: AI-Driven Post-Quantum DePIN Shield

The First Decentralized AI-PQC Middleware for EVM Ecosystems & Cross-Chain Bridges

## The "Quantum Collapse" & HNDL Threat

Legacy ECDSA rollups face an existential threat from quantum vectors. State actors and malicious entities are already executing "Harvest Now, Decrypt Later" (HNDL) attacks—archiving encrypted L1/L2 state data to decrypt retroactively once Cryptographically Relevant Quantum Computers (CRQCs) come online. The choice is binary: natively proof your infrastructure, or leave your ecosystem vulnerable to total TVL extraction.

## Core Architecture: The "Zero L1 Impact" Paradigm

QUBEX Sentinel operates as a non-invasive interceptor layer between the Sequencer and the Batcher, securing L2s against "Harvest Now, Decrypt Later" threats without modifying existing consensus logic. 
Integrating PQC directly on-chain forces unsustainable gas inflation and state bloat. QUBEX solves this via a decoupled pre-batcher execution:

1. AI Threat Detection (Chaos Engine): The middleware intercepts transactions, passing them through our predictive AI engine to detect anomalous "Harvest Now, Decrypt Later" sequencing patterns in real-time.
2. Decentralized Verification (Sentinel Nodes): Clean traffic is routed to our decentralized node network, where NIST-standard ML-DSA signatures are mathematically validated in ~0.17ms.
3. Compress: Only cryptographically verified, quantum-resistant payloads reach the op-batcher.

Result: Absolute PQC integrity on L2 with zero bytes of overhead added to L1 calldata. We enable post-quantum security at the speed of legacy ECDSA.

## Live Devnet Omnichain Benchmark (Concurrent Routing)
- Security Standard: NIST Level 5 (ML-DSA-87).
- Target Networks: 8 Tier-1 L2s (Base, Arbitrum, Optimism, Scroll, Mantle, BNB, Polygon, ETH).
- Execution: Concurrent Mempool Injection via Goroutines.
- Validation Latency: Sub-millisecond (~170µs - 650µs CPU cycle variance).
- L1 Gas Bloat: 0% footprint.

## Visual Proof

<img width="1152" height="815" alt="image" src="https://github.com/user-attachments/assets/23aaf3db-b918-45b0-aca3-ba30239f187b" />

<img width="1132" height="487" alt="image" src="https://github.com/user-attachments/assets/f45d4e84-bdb2-4059-9842-57c665273aa1" />


## Omnichain Reproduction & Local Verification
"Don't trust, verify." QUBEX Sentinel operates with absolute cryptographic transparency. Enterprise architects and protocol core developers can independently reproduce the stress test locally.

### Prerequisites
- Go Environment: Version 1.20 or higher.
- A burner EVM wallet with testnet gas across the 8 networks.
- Hardware: Multi-core CPU. *(Warning: The benchmark is designed to force 100% CPU thread exhaustion to simulate peak enterprise sequencer loads. Proceed with caution on commercial laptops)*.

### Execution Protocol
`bash
git clone [https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git](https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git)
cd Qubex-PQC-Benchmarks/omnichain-router
go mod tidy

# Inject your burner wallet's private key (NO 0x prefix)
# Linux/Mac:
export PRIVATE_KEY="YOUR_BURNER_PRIVATE_KEY"
# Windows (PowerShell):
# $env:PRIVATE_KEY="YOUR_BURNER_PRIVATE_KEY"

# Ignite the Devnet Broadcaster
go run tx.go

## Integration & Genesis Pilots

QUBEX Sentinel is currently in the MVP Phase, actively selecting high-TVL Cross-Chain Bridges, Oracles, and Tier-1 L2 Rollups for exclusive Genesis Pilot integrations

Networks participating in the Genesis Pilot phase are eligible for Strategic Integration Bounties (subsidized from our Ecosystem Grants allocation) to eliminate the financial friction of transitioning to a quantum-safe architecture.

### Initiate Integration:

To request a technical handoff, review the Trust Index framework, and secure a Genesis Pilot slot, submit your protocol's architecture for review by our core team.

Institutional Contact: spyridongagr@qubexsentinel.com

X (Twitter): @QUBEX_SENTINEL

[Submit Genesis Pilot Request Here](https://forms.gle/hmUdBiQz3PT2x8TT7)
