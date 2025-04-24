# Challenge 15

## Difficulty

Hard

## Title

fARMed

## Description

Ben’s router — a high-tech Plumber-issued device powered by an ARM chip — has been acting strange. Every time he tries to change the root password, it accepts the command, but the password never actually changes.

Suspicious, Gwen took a snapshot of the device's firmware and said,

> “This isn’t just bad code… this is deliberate sabotage.”

Looks like someone hardcoded a backdoor. Maybe Vilgax, maybe someone closer.

Ben needs root access to stop the malware beacon before it pings off-world.
Can you reverse the firmware and figure out what the actual root password is?

## Files

- `fARMware.bin`

## Solution

Extract the firmware using `binwalk` and navigate to `/usr/share/rpcd/ucode` at line 401.

## Flag

```text
cyn0x{r00t_4cce55_byp4553d_thr0ugh_f1rmw4re_h4rdc0d3}
```
