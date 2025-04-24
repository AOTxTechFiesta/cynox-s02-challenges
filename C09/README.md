# Challenge 9

## Difficulty

Hard

## Title

MaxSec

## Description

Grandpa Max has intercepted a suspicious executable from Vilgax’s secret transmission. It looks like a basic program — just a password box — but he knows there’s more to it.
_“I’ve seen this kind of thing before,” he mutters. “Coils of logic, hidden beneath the surface… like a snake waiting to strike.”_
Vilgax never makes it easy. Somewhere in this binary, the secret lies dormant. All it takes is the right password to wake it.

## Files

- `maxsec.exe` - The binary file for Windows
- `maxsec.out` - The binary file for Linux

## Solution

The binary is packed with [PyInstaller](https://pyinstaller.org/). We can extract it to get the compiled python code in `.pyc` format. Then we can reverse engineer / decompile the compiled code to find the flag.

Use [pyinstxtractor](https://github.com/extremecoders-re/pyinstxtractor/blob/master/pyinstxtractor.py) to extract the contents of the `maxsec.exe` file and then use [PyLingual](https://pylingual.io/) or [uncompyle6](https://pypi.org/project/uncompyle6/) to decompile the `.pyc` files to get the original python code.

## Flag

```text
cyn0x{gr4ndpa_m4x_d3bugged_th3_d4ng3r}
```
