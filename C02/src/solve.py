import base64

encoded_flag = open("secret.txt", "r").read()
decoded_flag = base64.b64decode(encoded_flag).decode()
print(decoded_flag)