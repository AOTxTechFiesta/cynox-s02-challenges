import base64

flag = "cyn0x{gwen_us3d_m4g1c_but_y0u_n33d_l0g1c_t0_unw1nd_th3_tw1st3d_c0d3}"

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

def encode_flag(flag):
    step1 = base64.b64encode(flag.encode()).decode()
    step2 = rot_n(step1, -69)
    step3 = base64.b64encode(step2.encode()).decode()
    step4 = rot_n(step3, -11)
    step5 = base64.b64encode(step4.encode()).decode()
    final = rot_n(step5, -9)
    return final

encoded = encode_flag(flag)
open ("secret.txt", "w").write(encoded)
print("Encoded flag written to secret.txt")
