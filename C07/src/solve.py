import re

BINARY_FILE = 'bin.out'

with open(BINARY_FILE, 'rb') as f:
    binary = f.read()

match = re.search(rb'cyn0x\{.*?\}', binary)
if match:
    flag = match.group(0).decode('utf-8')
    print(flag)
else:
    print("Flag not found")
