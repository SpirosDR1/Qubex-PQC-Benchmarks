# QUBEX SENTINEL: Omnichain Chaos Engine (Mainnet Infrastructure)

The first Decentralized AI-PQC Middleware for L1 Settlement Layers and EVM Rollups. 

This repository contains the production-ready concurrent routing infrastructure benchmarks for the QUBEX Sentinel network.

## The Architecture: Decoupled PQC Delivery

Current L1s and L2s face an existential threat from Harvest Now, Decrypt Later (HNDL) vectors. Natively upgrading consensus algorithms to quantum-safe standards is slow and causes catastrophic L1 state bloat. 

### QUBEX solves this via a decoupled pre-batcher layer:

1. The Routing Engine (LIVE): Sub-second, concurrent mempool piercing across multiple EVM chains simultaneously. 0% L1 gas bloat.

2. The Cryptographic Payload (PILOT PHASE): Delivering decentralized NIST-standard ML-DSA validations through our routing infrastructure before the transaction reaches the sequencer.

## Live Mainnet Benchmarks

We have successfully deployed and stress-tested the routing layer across 6 Tier-1 EVM Mainnets on enterprise-grade bare-metal Hetzner clusters. 

<img width="1092" height="807" alt="image" src="https://github.com/user-attachments/assets/a4150bc3-9bef-4d2c-bff9-92f963c6a064" />

<img width="977" height="197" alt="image" src="https://github.com/user-attachments/assets/01617de3-6a03-43e6-8721-5d7695baa007" />

> *Execution Mode: Concurrent Goroutines, dynamic gas logic with a 20% inclusion bump. Absolute redundancy verified.*

## Verification: "Don't Trust, Verify"
Enterprise architects and protocol core developers can independently reproduce our cross-chain routing latency and gas logic locally. 

### Execution Protocol:

```go
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
```

```go
cd Qubex-PQC-Benchmarks/omnichain-router
```

```go
go mod init qubex-pqc && go mod tidy
```

#### Phase 1: Cryptographic Payload Execution (PQC)

Validate the ML-DSA-87 signature integrity and nanosecond latency logic locally.

```go
go run omnichain-pqc-stress-test.go
```
#### Phase 2: Live Mainnet Routing Validation

*(Warning: Use a burner EVM wallet funded with gas for the target mainnets: Optimism($ETH),Polygon($POL),BSC($BNB),Ethereum($ETH),Base($ETH),Arbitrum($ETH)).*

*[INFRASTRUCTURE NOTE]: This benchmark fires concurrent goroutines against free, public EVM endpoints. Public RPCs impose strict rate-limits and may occasionally drop connections, resulting in a FAILED (RPC THROTTLED) or (GAS ERROR) log. This exposes the fragility of legacy public infrastructure, not the QUBEX routing logic. Our production Chaos Engine completely bypasses this by utilizing dedicated, enterprise-grade node RPCs.*

Validate the sub-second concurrent mempool piercing and 0% L1 gas bloat on live mainnets.

*For Mac/Linux:*

```go
export PRIVATE_KEY="YOUR_BURNER_PK"
```

*For Windows (CMD):*

```go
set PRIVATE_KEY=YOUR_BURNER_PK
```

```go
go run tx.go
```

## Genesis Integration Pilots

With the infrastructure routing architecture validated on Mainnet, QUBEX is opening limited slots for our 90-Day Zero-Cost Genesis Pilots.
We are actively onboarding high-TVL Rollups and RaaS platforms to integrate our full ML-DSA cryptographic payload via this exact routing engine.

* Institutional Contact: spyridongagr@qubexsentinel.com

* X (Twitter): [@QUBEX_SENTINEL](https://twitter.com/QUBEX_SENTINEL)

* Application: [Submit Genesis Pilot Request Here](https://forms.gle/hmUdBiQz3PT2x8TT7)
