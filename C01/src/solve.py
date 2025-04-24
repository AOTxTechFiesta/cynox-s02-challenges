secret = open("secret.txt", "r").read().split()
ascii_values = [int(i) for i in secret]

print("".join([chr(i) for i in ascii_values]))