# Challenge 7

## Title

Vilgax Returns

## Description

Vilgax embedded a secret deep within a rusty piece of Galvan tech. It asks for a password to unlock — but he’s cocky. He left the secret right there in the binary. Can you find it and prove you’re faster than him?

## Files

- `bin.exe` - The binary file for Windows
- `bin.out` - The binary file for Linux

## Solution

Use `strings` or a hex editor to find the password in the binary file.

With Linux, you can use the following command to extract strings from the binary file:

```bash
strings bin.out | grep -i cyn0x
```

With Windows, you can use the following command:

```bash
strings bin.exe | findstr cyn0x
```

## Flag

```text
cyn0x{v1lgax_und3r3stim4t3d_th3_p0w3r_0f_y0ur_t00ls}
```
