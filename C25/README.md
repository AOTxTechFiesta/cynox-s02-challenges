# Challenge 25

## Difficulty

Hard

## Title

Encoded in DNAliens

## Description

Ben has intercepted yet another of Vilgax’s apps — but this one is different. It’s been engineered with alien-level protection, deeply embedded, far beyond the reach of surface-level analysis.

Nothing useful shows up in plain text, and even the most sophisticated scanners come up empty. But Ben suspects that the secret — possibly a key to decoding Vilgax’s next attack — is locked away inside, masked from the average eye.

Only someone with the skills to dig into the alien tech can reveal what Vilgax tried so hard to hide.

Can you crack open the code?

## Files

- `app.apk`: The Android app file that contains the code and resources for the challenge.

## Solution

Use the following Frida script to dump the flag from the native code of the app.

```typescript
Java.perform(function() {
  const FlagClass = Java.use("dev.itsyourap.cynox.C28421");
  const object = FlagClass.$new();
  const flag = object.Ey1vO4MRjn379HD0RTYrCfHQQZ2dK4();
  console.log("[+] Flag: " + flag);
});
```

If you have a rooted device, you can run this script using Frida with the following command:

```bash
frida -U -f dev.itsyourap.cynox -l script.js --no-pause
```

This will attach to the app and execute the script, printing the flag to the console.

If you don't have a rooted device, you can either use a rooted emulator or use Frida's `frida-gadget` to inject the script into the app.

## Flag

```text
cyn0x{n4t1v3_r3v_15_4_p4th_t0_4l13n_1nt3ll1g3nc3}
```
