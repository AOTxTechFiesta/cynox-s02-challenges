flag = "cyn0x{r3c0v3ry_c0mpl3t3_s1gnal_p4tt3rn_d3crypt3d_b3tw33n_th3_l1n3s}"

def split(flag):
    ascii_values = [ord(char) for char in flag]
    set1, set2, set3 = [], [], []

    for value in ascii_values:
        part1 = value // 100
        part2 = (value // 10) % 10
        part3 = value % 10
        set1.append(part1)
        set2.append(part2)
        set3.append(part3)

    open("secret1.txt", "w").write(" ".join(map(str, set1)))
    open("secret2.txt", "w").write(" ".join(map(str, set2)))
    open("secret3.txt", "w").write(" ".join(map(str, set3)))

split(flag)
