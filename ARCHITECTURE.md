# Qubex Sentinel — Architecture

Post-quantum verification for assets on existing chains, with no
migration. This document describes both what runs today and the
architecture it is being built toward. Each component is labelled so the
two are never confused.

## Status legend

- **[BUILT]** — exists and runs in this repository today
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

What they need is a verification layer that adds post-quantum assurance
to their existing ECDSA flows, on their existing chains, that they can
audit and present to a regulator. That is the layer Qubex is building.

---

## 1. Cryptographic core  [BUILT]

ML-DSA-87 (NIST FIPS 204, level 5) signing and verification, benchmarked
for latency and correctness. This is the only component that runs today.
Everything below is built on this primitive.

## 2. Validation API  [DESIGN]

A stateless service that accepts an ML-DSA-87 public key, message, and
signature, verifies it, and returns a signed attestation of the result.

The design goals:

- Stateless and horizontally scalable — verification is independent per
  request.
- No custody of client funds or keys at any point.
- An interface a custodian can call as an added check over their
  existing ECDSA signing flow, without changing how they hold assets.

## 3. Institutional integration  [DESIGN]

Custodians call the Validation API as a post-quantum verification step
alongside their current process. No chain migration, no wallet rotation
for end clients, no smart-contract rewrites on the chains they use.

The integration is meant to be additive: it sits beside existing
infrastructure rather than replacing any part of it.

## 4. Open questions  [RESEARCH]

- **Attestation format.** A plain signed verification result, or a
  zero-knowledge proof that verification was performed correctly without
  revealing inputs.
- **Trust model.** How a custodian confirms the API returned an honest
  result — independent re-verification, on-chain anchoring of
  attestations, or both.
- **Anchoring.** Whether verification results need to be committed
  on-chain at all, and at what cost/benefit.

---

## Design principles

- **No migration.** Security is added where the assets already are.
- **NIST-standard only.** ML-DSA (FIPS 204), so the cryptography is one a
  regulated institution can stand behind.
- **No custody.** Qubex verifies; it never holds keys or funds.
- **Auditable.** Every claim in this document maps to a status label.
  What is built can be run. What is designed is marked as designed.
