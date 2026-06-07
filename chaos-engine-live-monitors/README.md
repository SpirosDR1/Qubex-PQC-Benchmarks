# QUBEX Sentinel: Chaos Engine (Live RPC Monitors)

## Overview
The modular scaling ecosystem is actively compounding migration debt. The scripts in this repository are part of the QUBEX Chaos Engine, designed to publicly benchmark legacy cryptographic latency against quantum-resistant primitives (NIST ML-DSA). 

By hooking directly into live public RPC endpoints, these monitors expose the critical vulnerability windows within EVM sequencers, cross-chain bridges, and Rollup-as-a-Service (RaaS) frameworks.

The math is absolute. We do not generate these numbers; your own infrastructure reports them.

## The Vulnerability Matrix (HNDL Exposure)
When a sequencer or validator requires X milliseconds to batch and finalize a transaction, the payload sits unencrypted in the mempool. Against a simulated quantum HNDL vector (approx. 800ms), any latency exceeding this threshold creates a critical Exposure Window where ECDSA signatures can be extracted and payloads manipulated before block inclusion.

## Verification Guide: Run the Audit
You do not need to trust our benchmarks. Verify your network's exposure locally. These scripts are agnostic and designed to audit any EVM-compatible endpoint.

### 1. Prerequisites
Ensure you have Python 3 installed. Install the required dependencies:

```bash
pip install web3 colorama
```

### 2. Execution
Run the target-specific monitor from your terminal.

Example:

```bash
python chaos_arbitrum_monitor.py
```

```bash
python chaos_raas_monitor.py
```

*(Note: To test a custom infrastructure—L1, L2, L3, or Bridge—simply modify the RPC_URL variable inside any of the scripts to your target endpoint).*

### 3. Understanding the Telemetry

Upon execution, the Chaos Engine reads live block headers directly from the target RPC.

[CRITICAL ALERT] (RED): The sequencer latency exceeds the quantum attack threshold. The output displays the exact millisecond window your TVL is exposed.

QUBEX Patch Ready (GREEN): Displays the QUBEX Sentinel mitigation benchmark. Our decoupled routing engine processes ML-DSA validations off-chain in nanoseconds, with 0% L1 gas bloat, securing the payload before the EVM state is broadcasted.

## The QUBEX Standard

Vulnerability-as-a-Service is not a roadmap. Waiting for multi-year L1 hard forks is a failure of execution.
QUBEX Sentinel injects sub-millisecond post-quantum validation directly into production stacks today.
Secure the routing.
