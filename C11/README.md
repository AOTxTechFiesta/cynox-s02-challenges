# Challenge 11

## Difficulty

Easy

## Title

Snapshot Secrets

## Description

Gwen snapped this photo of Ben during one of their field missions — nothing seemed unusual at first glance. But she had a hunch.

> “Photos don’t just capture images,” she said, “sometimes they store… secrets.”

Can you uncover what Gwen sensed?

## Files

- `image.jpg`: An image of Ben

## Solution

Just extract the metadata from the image using `exiftool`:

```bash
exiftool image.jpg
```

This will show you the metadata of the image, including a comment field that contains the flag.

## Flag

```text
cyn0x{n0t_3v3ry_p1x3l_m4tt3rs_s0m3t1m3s_1ts_4ll_4b0ut_th3_c4pt10n_y0u_c4nt_s33}
```
