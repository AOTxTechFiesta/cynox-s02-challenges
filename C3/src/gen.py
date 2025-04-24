import random

flag = "cyn0x{x0r3d_truth_fr0m_dual_streams_0f_galvan_crypt0_l0ck}"

def split_and_xor(flag):
    ascii_values = [ord(char) for char in flag]
    set1, set2 = [], []

    for value in ascii_values:
        part1 = random.randint(10**8, 10**15 - 1)
        part2 = value ^ part1
        set1.append(part1)
        set2.append(part2)

    open("secret1.txt", "w").write(" ".join(map(str, set1)))
    open("secret2.txt", "w").write(" ".join(map(str, set2)))

split_and_xor(flag)
