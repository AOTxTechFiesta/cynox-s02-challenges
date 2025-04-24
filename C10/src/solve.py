import base64

def rot_n(text, n):
    result = []
    for char in text:
        if 'a' <= char <= 'z':
            result.append(chr((ord(char) - ord('a') + n) % 26 + ord('a')))
        elif 'A' <= char <= 'Z':
            result.append(chr((ord(char) - ord('A') + n) % 26 + ord('A')))
        else:
            result.append(char)
    return ''.join(result)

def decode_flag(flag):
    step1 = rot_n(flag, 9)
    step2 = base64.b64decode(step1).decode()
    step3 = rot_n(step2, 11)
    step4 = base64.b64decode(step3).decode()
    step5 = rot_n(step4, 69)
    final = base64.b64decode(step5).decode()
    return final

encoded = open("secret.txt", "r").read()
decoded = decode_flag(encoded)
print(decoded)
