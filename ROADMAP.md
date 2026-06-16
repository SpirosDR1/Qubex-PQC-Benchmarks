# Qubex Sentinel — Roadmap

This file tracks what exists versus what is planned. Honesty about
status is the point: anything not marked "Built" does not run yet.

## Built

- ML-DSA-87 benchmark — sign/verify latency and correctness
  measurement (FIPS 204, level 5). Runs today in this repo.

- Verification API — a stateless /verify endpoint that accepts an
  ML-DSA-87 public key, message, and signature and returns whether
  it verifies. Holds no user keys, stores nothing. Runs today.

- Signed attestation — every verification result is signed with the
  service's own ML-DSA-87 key and returned as a portable receipt,
  verifiable against the public key at /attestation-key. Runs today.

- Persistent attestation key — the attestation key is loaded from a
  fixed secret and is stable across restarts, giving the service a
  permanent cryptographic identity that receipts can be pinned to.
  Runs today.

## Planned

- Institutional integration — calling the API as an added
  post-quantum check over existing ECDSA custody flows, no migration.

- Security review — the near-term path is peer and community review:
  direct engagement with cryptographers working on ML-DSA and the
  underlying library, plus participation in the NIST PQC community.
  A paid third-party audit becomes the right move once there's a
  funded pilot or real institutional stake to justify it — not
  before.

## Research

- ZK proof of correct verification — the current attestation signs
  the result the service returned; it does not prove the verification
  itself was performed correctly. A zero-knowledge proof would close
  that trust gap.

- On-chain anchoring of attestations: needed or not?

_Last updated: 16 June 2026_
