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
directly against official, NIST-aligned known-answer test vectors from
github.com/post-quantum-cryptography/KAT (pure mode). Unlike
cross_check.go, this calls circl directly rather than through the API,
because the KAT vectors use a non-empty 16-byte context string, which
the current /verify endpoint does not expose.

Tested against both official signing modes:

```bash
go run kat_check.go -file kat_MLDSA_87_det_pure.rsp
go run kat_check.go -file kat_MLDSA_87_hedged_pure.rsp
```

Last run result: 200/200 official vectors verified correctly — 100/100
deterministic mode, 100/100 hedged mode (the mode the live API actually
uses in production).

<img width="622" height="182" alt="image" src="https://github.com/user-attachments/assets/0ef19dbe-91a5-48e8-8602-b7ef51e4a1eb" />

<img width="622" height="190" alt="image" src="https://github.com/user-attachments/assets/c1712d32-7618-4e26-be85-073c3d005f2c" />

## What this does and doesn't prove

Together, these two checks confirm: (1) the API's bridge code correctly
wires hex input to the cryptographic library across genuine and
adversarial cases, and (2) the underlying ML-DSA-87 implementation
itself produces correct results against official, independently-known
test data, in both signing modes — not just our own self-generated test
cases.

What's still open: this validates the cryptographic correctness of
verification. It does not prove that a deployed instance of the API is
running unmodified, audited code at any given moment (that's what a
TEE-based attestation, on the roadmap, would address) — see ROADMAP.md
and ARCHITECTURE.md for what's built versus still open.
