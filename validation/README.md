# Independent Validation

This directory contains standalone checks that do not rely on trusting
the API's own claims.

## Integration cross-check

`cross_check.go` calls cloudflare/circl's ML-DSA-87 Verify function
directly, bypassing the API's HTTP layer entirely, and compares the
result against what the live API says for the same inputs — across a
genuine signature, a tampered signature, a wrong message, and a
signature from the wrong key.

Run it yourself:

```bash
go run cross_check.go
```

This tests the integration code (the part that converts hex input to
bytes and calls the cryptographic library) — not the underlying
ML-DSA-87 math itself, which is cloudflare/circl's responsibility and
already runs in Cloudflare's own production infrastructure.

Last run result: all 4 cases passed (genuine ✓, tampered ✓, wrong
message ✓, wrong key ✓).

## Official NIST KAT validation

`kat_check.go` validates cloudflare/circl's ML-DSA-87 implementation
directly against 100 official, NIST-aligned known-answer test vectors
from github.com/post-quantum-cryptography/KAT (deterministic, pure
mode). Unlike cross_check.go, this calls circl directly rather than
through the API, because the KAT vectors use a non-empty 16-byte
context string, which the current /verify endpoint does not expose.

Run it yourself:

```bash
go run kat_check.go -file kat_MLDSA_87_det_pure.rsp
```

Last run result: 100/100 official vectors verified correctly.

## What this does and doesn't prove

Together, these two checks confirm: (1) the API's bridge code correctly
wires hex input to the cryptographic library across genuine and
adversarial cases, and (2) the underlying ML-DSA-87 implementation
itself produces correct results against official, independently-known
test data — not just our own self-generated test cases.

What's still open: this validates the cryptographic correctness of
verification. It does not prove that a deployed instance of the API is
running unmodified, audited code at any given moment (that's what a
TEE-based attestation, on the roadmap, would address) — see ROADMAP.md
and ARCHITECTURE.md for what's built versus still open.
