# Qubex Genesis Framework — Integration Path for Custodians & Infrastructure

## Purpose

This document describes the path by which an institutional custodian or
infrastructure operator can adopt post-quantum verification through
Qubex Sentinel. It is a blueprint for where the integration is going. It
distinguishes, at every step, what exists today from what is being
built.

## Where things stand

Built today:
- The ML-DSA-87 cryptographic core (FIPS 204, level 5), benchmarked for
  performance and correctness.
- The Verification API — a live, stateless service that verifies an
  ML-DSA-87 signature and returns the result with a signed attestation,
  verifiable against a public key at /attestation-key. Early, but real.

In development: the integration path described below (Phases 1–3). The
API exists; the institutional integration path has not yet been run with
a partner. We say so plainly, because a partner deciding whether to build
on Qubex needs to know exactly what is real.

## The integration path

### Phase 1 — Verification pilot
- **Goal:** confirm, against the partner's real traffic patterns, that
  post-quantum verification meets their latency and throughput needs.
- **How:** the partner sends verification requests (public key, message,
  signature) to the Qubex Verification API in parallel with their
  existing flow. Nothing in their production path changes; Qubex runs
  alongside as an observer.
- **Outcome:** measured latency and throughput data, and an honest
  assessment of fit — including where it doesn't fit yet.

### Phase 2 — Active verification layer
- **Goal:** post-quantum assurance on the partner's highest-value
  operations.
- **How:** critical transactions are verified through Qubex before the
  partner's system finalizes them. ML-DSA-87 verification becomes a
  required check on the assets that matter most.
- **Outcome:** the partner's most sensitive value carries a
  NIST-standard post-quantum verification step, with no change to the
  chains they use or the keys their clients hold.

### Phase 3 — Standard integration
- **Goal:** verification as a permanent, native part of the partner's
  process.
- **How:** the Qubex verification step is built directly into the
  partner's pipeline rather than called as an external service.
- **Outcome:** post-quantum verification becomes a default property of
  the partner's infrastructure, not an add-on.

## What the partner keeps

- Their chains. No migration.
- Their custody model. Qubex holds no keys and no funds.
- Their clients' wallets. No rotation required.
- Their execution speed. Verification runs alongside, not inside, the
  critical path until the partner chooses to enshrine it.

## Pilot slots

We are preparing a limited number of Phase 1 verification pilots. If you
operate institutional custody or rollup infrastructure and want to
profile post-quantum verification against your own traffic, reach out.

- Email: spyridongagr@qubexsentinel.com
- X: [@QUBEX_SENTINEL](https://x.com/QUBEX_SENTINEL)
- Apply: [Genesis Pilot Request](https://forms.gle/hmUdBiQz3PT2x8TT7)
