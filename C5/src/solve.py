secret = open("secret.txt", "r").read().strip()
def rot27(text):
    result = ''
    for char in text:
        if char.isalpha():
            base = ord('A') if char.isupper() else ord('a')
            # Rotate backwards by 1 (equivalent to ROT27)
            rotated = chr((ord(char) - base + 1) % 26 + base)
            result += rotated
        else:
            result += char  # Leave non-alpha characters unchanged
    return result

print(rot27(secret))