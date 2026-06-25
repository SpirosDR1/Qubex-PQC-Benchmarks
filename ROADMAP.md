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
- Official NIST KAT validation — 200/200 official known-answer test
  vectors verified correctly: 100/100 deterministic mode,
  100/100 hedged mode (the mode the live API uses in production).
  Runs today via /validation/kat_check.go.
- Load test — 2,100 requests across four concurrency levels (10, 50,
  100, 200 simultaneous), zero failures, zero incorrect verifications.
  Throughput plateau ~110 req/sec on current hosting. Runs today via
  /loadtest.go.
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
  that trust gap. This is an open research problem — lattice
  operations inside ZK circuits are computationally expensive and
  not yet solved at production scale by anyone.
- Hardware-attested verification (TEE) — running the verification
  inside a hardware enclave (e.g., Intel SGX, AWS Nitro Enclaves)
  would allow clients to verify correct execution without trusting
  the service operator. This eliminates the current trust assumption
  without requiring a ZK proof, but raises its own enclave
  attestation management questions. A viable intermediate step
  between the current model and a full ZK proof.
- On-chain anchoring of attestations: needed or not?
_Last updated: 25 June 2026_
