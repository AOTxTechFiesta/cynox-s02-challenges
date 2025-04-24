# Challenge 6

## Difficulty

Medium

## Title

Vilgaxian Pressure Seal

## Description

Intercepted from the remnants of a failed warp jump near a Vilgaxian war cruiser, this data segment was sealed under intense system pressure — a technique believed to distort and condense information beyond conventional limits.

Our techs say the payload was layered with misdirection to avoid interception. Whatever it is, it’s locked inside a pressurized data vault.

The file `secret.txt` holds the full sequence. Can you relieve the pressure and retrieve the hidden intel?

## Files

- `secret.txt` - The encoded message

## Solution

Base64 decode the string in `secret.txt` and then decompress it using `lzma`.

```bash
base64 -d secret.txt | lzma -d
```

## Flag

```text
cyn0x{v1lgax_d4ta_pr3ssur3_s3al_r3l3as3d_byt3s_r3st0r3d}
```
