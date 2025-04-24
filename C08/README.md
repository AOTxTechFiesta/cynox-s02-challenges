# Challenge 8

## Difficulty

Medium

## Title

Vilgax Cipher

## Description

Vilgax knows his last attempt failed, so he’s gone full stealth mode.

He embedded the flag in a scrambled form inside a mysterious binary recovered from the Omnitrix core. The interface looks harmless, but deep within, the truth awaits.

Can you find the flag?

## Files

- `bin.exe` - The binary file for Windows
- `bin.out` - The binary file for Linux

## Solution

Use Ghidra to analyze the binary. The main function initializes a buffer with random hex values and then xor's them with the XOR_KEY `0x42` which results in the flag.

## Flag

```text
cyn0x{v1lgax_3nc0d3d_y0u_must_r3v3rs3_1t}
```
