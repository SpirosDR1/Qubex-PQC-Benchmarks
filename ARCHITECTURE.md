# Qubex Sentinel — Architecture

Post-quantum verification for assets on existing chains, with no
migration. This document describes both what runs today and the
architecture it is being built toward. Each component is labelled so the
two are never confused.

## Status legend

- **[BUILT]** — exists and runs today
- **[DESIGN]** — intended architecture, not yet implemented
- **[RESEARCH]** — open question, not yet decided

---

## The problem

Institutional crypto sits on chains secured by ECDSA. A sufficiently
powerful quantum computer can derive private keys from the public keys
already visible on-chain. Institutions cannot respond the way retail can:

- They can't migrate to a new "quantum-safe" chain — custody
  relationships, prime-broker agreements, and regulatory frameworks are
  bound to the chains they're on.
- They can't wait for base-layer consensus to upgrade — Bitcoin's PQ
  signature scheme is pre-consensus, and Ethereum's path spans multiple
  hard forks across years.

What they need is a verification checkpoint inside their own existing
authorization flows — NIST-standard, independently checkable, auditable
and presentable to a regulator. That is the layer Qubex is building.

Honest limit, stated plainly: this does not protect an already-exposed
public key from an attacker broadcasting a forged signature directly to
the chain. No external verification layer can — that is a base-layer
protocol problem, closed only by chain migration or opt-in account
abstraction (e.g. Ethereum's EIP-8141 direction, still in proposal
stage as of mid-2026). What Qubex adds is narrower and real: a
verification checkpoint inside the institution's own process, and
demonstrable readiness for when that protocol-level path matures.

---

## 1. Cryptographic core  [BUILT]

ML-DSA-87 (NIST FIPS 204, level 5) signing and verification, benchmarked
for latency and correctness. This is the primitive everything below is
built on.

## 2. Verification API  [BUILT]

A stateless service that accepts an ML-DSA-87 public key, message, and
signature, verifies it, and returns the result together with a signed
attestation. Holds no user keys and stores nothing. Runs today.

- Stateless and horizontally scalable — verification is independent per
  request.
- No custody of client funds or keys at any point.
- Each result is signed with the service's own ML-DSA-87 attestation
  key and is verifiable against the public key at /attestation-key,
  making the result a portable receipt.

Honest limit of what runs today: the attestation signs *the result the
service returned* — it does not yet prove the verification was performed
correctly.

## 3. Persistent attestation key  [BUILT]

The attestation key is loaded from a fixed secret at startup and is
stable across restarts. This gives the service a permanent cryptographic
identity: a receipt issued today verifies against the same public key
later, so receipts can be pinned. The private key is held only as a
deployment secret and never committed to source.

## 4. Institutional integration  [DESIGN]

Custodians call the Verification API as a post-quantum verification step
alongside their current process. No chain migration, no wallet rotation
for end clients, no smart-contract rewrites on the chains they use.

The integration is additive: it sits beside existing infrastructure
rather than replacing any part of it.

## 5. Open questions  [RESEARCH]

- **ZK proof of correct verification.** The current attestation signs
  the returned result; it does not prove the verification itself ran
  correctly. A zero-knowledge proof would close that trust gap without
  revealing inputs. This is an open research problem — lattice
  operations inside ZK circuits are computationally expensive and not
  yet solved at production scale by anyone.

- **Hardware-attested execution (TEE).** Running the verification inside
  a hardware enclave (e.g., Intel SGX, AWS Nitro Enclaves) would allow
  clients to verify correct execution without trusting the service
  operator — eliminating the current trust assumption without requiring
  a ZK proof. This raises its own enclave attestation management
  questions and is a viable intermediate step between the current model
  and a full ZK proof.

- **Trust model.** How a custodian confirms the API returned an honest
  result — independent re-verification, on-chain anchoring, TEE
  attestation, or some combination.

- **Anchoring.** Whether verification results need to be committed
  on-chain at all, and at what cost/benefit.

---

## Design principles

- **No migration.** Security is added where the assets already are.
- **NIST-standard only.** ML-DSA (FIPS 204), so the cryptography is one
  a regulated institution can stand behind.
- **No custody.** Qubex verifies; it never holds keys or funds.
- **Auditable.** Every claim in this document maps to a status label.
  What is built can be run. What is designed is marked as designed.
