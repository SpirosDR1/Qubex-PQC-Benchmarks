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
sign and verify performance, using the cloudflare/circl implementation.

This is the cryptographic primitive the rest of the system is built on.
It is a measurement tool. It does not secure any chain and is not a
production service — see ROADMAP for what is built versus designed.

## Run it

Requires Go 1.21+.

```go
git clone https://github.com/SpirosDR1/Qubex-PQC-Benchmarks.git
```

```go
cd Qubex-PQC-Benchmarks/benchmarks_code
```

```go
go mod tidy
```

```go
go run main.go
```

Output: average sign latency, average verify latency, and a validity
check, in nanoseconds. Numbers are hardware-dependent — report your CPU
when you share results.


<img width="495" height="206" alt="image" src="https://github.com/user-attachments/assets/b25143ae-46cc-4b9f-afe3-c581f0eac48c" />


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
| Validation API             | Design   |
| Institutional integration  | Design   |

## Contact

- Email: spyridongagr@qubexsentinel.com
- X: [@QUBEX_SENTINEL](https://x.com/QUBEX_SENTINEL)
- Apply: [Genesis Pilot Request](https://forms.gle/hmUdBiQz3PT2x8TT7)

## License

MIT
