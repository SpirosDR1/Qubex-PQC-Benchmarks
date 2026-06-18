# Independent Validation

This directory contains a standalone cross-check that does not rely on
trusting the API's own claims.

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
already runs in Cloudflare's own production infrastructure. Validating
against NIST's official ACVP known-answer test vectors is a further
step, not yet done — tracked on the roadmap.

Last run result: all 4 cases passed (genuine ✓, tampered ✓, wrong
message ✓, wrong key ✓).
