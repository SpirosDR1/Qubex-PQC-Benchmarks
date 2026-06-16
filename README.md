# Qubex Sentinel — PQC Benchmarks

Post-quantum signature validation for institutional crypto custody.
This repository contains the reference benchmark for the cryptographic
core: NIST ML-DSA-87 signing and verification.

## Why this exists

Institutions hold their assets on chains secured by ECDSA. Those
signatures are exposed to "Harvest Now, Decrypt Later" — public keys
visible on-chain today can be attacked the day a cryptographically
relevant quantum computer exists. Custodians can't migrate to new
chains: their positions, custody relationships, and compliance
frameworks are fixed where they are.

Qubex Sentinel is being built to give them post-quantum verification
for the assets they already hold, on the chains they already use, with
no migration. The standard that wins this will be the one institutions
can put in front of a regulator. That's what we're building toward.

## What's in this repo today

A reproducible benchmark of ML-DSA-87 (NIST FIPS 204, security level 5)
sign and verify performance, using the cloudflare/circl implementation,
plus a stateless verification API with signed attestations (see below).

This is the cryptographic primitive the rest of the system is built on.
The benchmark is a measurement tool. It does not secure any chain and is
not a production service — see ROADMAP for what is built versus designed.

## Run it

Requires Go 1.21+.

```bash
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
```
```bash
cd Qubex-PQC-Benchmarks/benchmarks_code
```
```bash
go mod tidy
```
```bash
go run main.go
```

Output: average sign latency, average verify latency, and a validity
check, in nanoseconds. Numbers are hardware-dependent — report your CPU
when you share results.

<img width="495" height="206" alt="image" src="https://github.com/user-attachments/assets/b25143ae-46cc-4b9f-afe3-c581f0eac48c" />

## Verification API

Beyond the benchmark, the repo includes a stateless verification
service. It exposes a /verify endpoint that checks an ML-DSA-87
signature against a public key and message, returns whether it is valid,
and returns a signed attestation of that result. It holds no user keys
and stores nothing.

Each verification result is signed with the service's own ML-DSA-87
attestation key, making the result a portable receipt that anyone can
verify independently against the published attestation public key
(/attestation-key). The attestation key is persistent and stable across
restarts, so receipts can be pinned to a fixed identity. Note: the
attestation signs the *result the service returned* — it does not yet
prove the verification was performed correctly. A zero-knowledge proof of
correct verification is on the roadmap.

```bash
cd Qubex-PQC-Benchmarks
```
```bash
go run ./api
```

*Then in another terminal:*

```bash
curl -s localhost:8080/demo > demo.json
```
```bash
curl -s -X POST localhost:8080/verify -H "Content-Type: application/json" -d @demo.json
```

Endpoints: `/` (info), `/health`, `/verify` (POST), `/attestation-key`,
`/demo`.

## Methodology

- Scheme: ML-DSA-87 (FIPS 204, level 5)
- Library: cloudflare/circl
- 1,000-iteration warm-up, then 100,000 timed iterations per operation
- Single-threaded wall-clock timing

Level 5 is the highest ML-DSA parameter set. Larger signatures, maximum
security margin — the conservative choice for custody, where the cost of
being wrong is the whole position.

## Status

| Component                  | State    |
|----------------------------|----------|
| ML-DSA-87 benchmark        | Built    |
| Verification API           | Built    |
| Signed attestation         | Built    |
| Persistent attestation key | Built    |
| ZK proof of verification   | Research |
| Institutional integration  | Design   |

## Contact

- Email: spyridongagr@qubexsentinel.com
- X: [@QUBEX_SENTINEL](https://x.com/QUBEX_SENTINEL)
- Apply: [Genesis Pilot Request](https://forms.gle/hmUdBiQz3PT2x8TT7)

## License

MIT
