# Challenge 13

## Difficulty

Medium

## Title

Maxed Out

## Description

Gwen found a crumpled letter stuffed into the dashboard of the Rustbucket — Grandpa Max's handwriting, worn by time and travel.
It looks like he left it for Ben, right before vanishing without a trace.
He warned not to go looking for him… but that’s never stopped you before.

Can you find out what Max was hiding, and where he might have gone?

## Files

- `max_letter.jpg`: The letter from Grandpa Max.

## Solution

Extract the flag from the audio using `steghide`. The passphrase is written in the letter.

```bash
steghide extract -sf max_letter.jpg -p TrustNoVilgax2025
```

## Flag

```text
cyn0x{null_v01d_3sc4p3_s3qu3nc3_1n1t14t3d_7r4ck_0ffgr1d}
```
