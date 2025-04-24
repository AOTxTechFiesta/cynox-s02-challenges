# Challenge 23

## Difficulty

Easy

## Title

Alien Signal Decoder

## Description

Ben intercepted a mysterious Android app from one of Vilgax’s secret transmissions. At first glance, it looks like a harmless tool — maybe even of no use. But Ben has a hunch... something’s buried deep inside.

Vilgax isn’t one to rely on flashy security. He trusts in complexity and layers. But he made one mistake — underestimating Ben.

Can you help Ben crack open the alien tech and reveal the message Vilgax tried to hide in the code?

## Files

- `app.apk`: The Android app file that contains the code and resources for the challenge.

## Solution

Decompile the app using `jadx` and get the flag in `C18201` class of `dev.itsyourap.cynox` package.

## Flag

```text
cyn0x{r3v3rs1ng_4l13n_c0d3_w1th_th3_0mn1tr1x}
```
