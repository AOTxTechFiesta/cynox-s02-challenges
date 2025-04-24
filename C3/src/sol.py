s1 = open("secret1.txt", "r").read().split()
x1 = [int(i) for i in s1]

s2 = open("secret2.txt", "r").read().split()
x2 = [int(i) for i in s2]

for i in range(len(x1)):
    x1[i] = chr(x1[i] ^ x2[i])

print(''.join(x1))