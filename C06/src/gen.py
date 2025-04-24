import lzma
import base64

flag = "cyn0x{v1lgax_d4ta_pr3ssur3_s3al_r3l3as3d_byt3s_r3st0r3d}"

# Compress the flag using LZMA (Alone format)
compressed_data = lzma.compress(flag.encode(), format=lzma.FORMAT_ALONE)

# Encode to Base64
encoded_data = base64.b64encode(compressed_data).decode()

# Write to file
with open("secret.txt", "w") as f:
    f.write(encoded_data)

print("Compressed and encoded flag written to secret.txt")
