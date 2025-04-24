import base64
import lzma

with open('secret.txt', 'r') as f:
    encoded = f.read().strip()

# Decode from Base64
compressed_data = base64.b64decode(encoded)

# Decompress using LZMA Alone (raw format)
decompressor = lzma.LZMADecompressor(format=lzma.FORMAT_ALONE)
flag = decompressor.decompress(compressed_data).decode()

print(flag)
