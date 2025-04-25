# Challenge 16

## Difficulty

Medium

## Title

Null Void Escape Protocol

## Description

Before heading into the Null Void, Grandpa Max planted a special flag — a beacon marking the exact position on Earth where he plans to return.

Time moves differently in the Null Void, and communication is nearly impossible. But he left behind one chance: a mysterious console connected to his Plumber tech.

To activate the return sequence, Ben needs to provide the precise location of the flag Grandpa Max left behind. The system isn’t very friendly — it only jumps to a position if you guess the coordinates exactly.

Gwen suspects there’s a hidden trick to it. Kevin says it’s a Plumber’s challenge only a real pro could solve.

Can you find the flag's position and bring Grandpa Max back to Earth?

Connect to the console using:

```bash
nc escape-protocol.cynox.aotfiesta25.tech 39213
```

## Files

- `escape_protocol.c`: C source code for the challenge.
- `escape_protocol`: Compiled binary for the challenge.

## Solution

We can calculate the address of the `main` and `plumber_warp_gate` functions using `objdump`. Then we need to calculate the offset of the `plumber_warp_gate` function from the `main` function.

After connecting to the console, we need to get the address of the `main` function in the binary. We can provide `%p` to the input to get the address of the `main` function. Then we can calculate the address of the `plumber_warp_gate` function using the offset we calculated earlier.

Finally, we can provide the address of the `plumber_warp_gate` function to the console to get the flag.

## Flag

```text
cyn0x{b3n_f0und_th3_p0rt4l_jump_p01nt_thr0ugh_pwn}
```
