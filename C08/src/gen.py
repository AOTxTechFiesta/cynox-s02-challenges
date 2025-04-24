flag = "cyn0x{v1lgax_3nc0d3d_y0u_must_r3v3rs3_1t}"

xor_key = 0x42

a = bytes([b ^ xor_key for b in flag.encode()])
# Print in hex array format
print("0x" + ", 0x".join([f"{b:02x}" for b in a]))