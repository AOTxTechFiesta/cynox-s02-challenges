# Challenge 10

## Difficulty

Medium

## Title

Arcane Layers

## Description

Gwen focused her magic, sensing residual energy lingering around an old Plumber safe. With a soft incantation and a flick of her charm, glowing numbers appeared on the lock: **9 11 69**.

She turned the dial in that exact sequence. Click. The safe opened.

Inside? No alien tech, no ancient relic — just a single piece of parchment filled with an unintelligible mess of symbols and characters.

> “This isn’t magical… it’s encoded,” Gwen said. “And I think the safe code — **9, 11, 69** — it’s more than just a key to the lock. It can also be used to decode the message somehow just like the safe.”

## Files

- `secret.txt`: The encoded message.

## Solution

ROT9 → Base64 decode → ROT11 → Base64 decode → ROT69 → Base64 decode

## Flag

```text
cyn0x{gwen_us3d_m4g1c_but_y0u_n33d_l0g1c_t0_unw1nd_th3_tw1st3d_c0d3}
```
