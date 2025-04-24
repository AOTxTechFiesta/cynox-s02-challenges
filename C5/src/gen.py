flag = "cyn0x{cl0s3_3n0ugh_f0r_a_galvan_r0tati0nal_scrambl3}"

def rot25(text):
    result = ''
    for char in text:
        if char.isalpha():
            base = ord('A') if char.isupper() else ord('a')
            # Rotate backwards by 1 (equivalent to ROT25)
            rotated = chr((ord(char) - base - 1) % 26 + base)
            result += rotated
        else:
            result += char  # Leave non-alpha characters unchanged
    return result

secret = rot25(flag)
open("secret.txt", "w").write(secret)