# Challenge 18

## Difficulty

Hard

## Title

Null Void Virus

## Description

The Null Void virus has struck again.

Ben tried booting up his Plumber-issued terminal, only to find all the system binaries completely wiped. No commands. No utilities. Just an eerie silence across the Omnitrix console.

Connect to the terminal using:

```bash
nc plumber-terminal.ctf.aotfiesta25.tech 41213
```

## Solution

```bash
cd                                                           # Go to ~
echo .*                                                      # List hidden files
while FSTREAM= read -r line; do echo "$line"; done < .flag   # Read the `.flag` file
```

## Flag

```text
cyn0x{t3rm1nal_0ffl1n3_n0_c0mm4nd5_y3t_fl4g_w45_4lw4y5_th3r3}
```
