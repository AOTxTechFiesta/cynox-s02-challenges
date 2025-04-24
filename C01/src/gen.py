flag = "cyn0x{transmission_dec0ded_fr0m_plumber_t3rm1nal_0verride_galvan_sig}"

s = ""

for i in flag:
    s += str(ord(i)) + " "

open("secret.txt", "w").write(s.strip())