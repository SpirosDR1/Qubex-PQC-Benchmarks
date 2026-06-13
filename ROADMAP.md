# Qubex Sentinel — Roadmap

This file tracks what exists versus what is planned. Honesty about
status is the point: anything not marked "Built" does not run yet.

## Built

- ML-DSA-87 benchmark — sign/verify latency and correctness
  measurement (FIPS 204, level 5). Runs today in this repo.
  
- Validation API — a stateless /verify endpoint that accepts an
  ML-DSA-87 public key, message, and signature and returns whether
  it verifies. Holds no keys, stores nothing. Runs today in this repo.

## Planned

- Attestation format — signed verification result, with research
  into a zero-knowledge proof variant.
- Institutional integration — calling the API as an added
  post-quantum check over existing ECDSA custody flows, no migration.

## Open questions

- On-chain anchoring of attestations: needed or not?
- Trust model: how a custodian independently confirms a verification
  result.

_Last updated: 14 June 2026_
