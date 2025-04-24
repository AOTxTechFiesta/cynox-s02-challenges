s1 = open("secret1.txt", "r").read().split()
s2 = open("secret2.txt", "r").read().split()
s3 = open("secret3.txt", "r").read().split()

ascii_values = []
for i in range(len(s1)):
    ascii_value = int(s1[i] + s2[i] + s3[i])
    ascii_values.append(ascii_value)

flag = ''.join([chr(value) for value in ascii_values])
print(flag)