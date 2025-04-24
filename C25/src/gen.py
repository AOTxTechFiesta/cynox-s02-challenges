import random

FLAG = "cyn0x{n4t1v3_r3v_15_4_p4th_t0_4l13n_1nt3ll1g3nc3}"


def generate_alphanumeric_string(length):
    characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    first_character = random.choice(
        "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    return first_character + ''.join(random.choice(characters) for _ in range(length - 1))


lst = []
dct = {}
allfn = []

for i in range(len(FLAG)):
    letter = FLAG[i]
    rand_num = random.randint(-10**5, 10**5)
    xor_key = ord(letter) ^ rand_num

    f1 = generate_alphanumeric_string(30)
    f2 = generate_alphanumeric_string(30)
    f3 = generate_alphanumeric_string(30)
    allfn.append(f"int {f1}();")
    allfn.append(f"int {f2}();")
    allfn.append(f"int {f3}();")

    fn1 = f"""
int {f1}() {{
  return {rand_num};
}}
"""

    fn2 = f"""
int {f2}() {{
  return {xor_key};
}}
"""

    fn3 = f"""
int {f3}()
 {{
  return {f1}() ^ {f2}();
}}
"""
    lst.append(fn1)
    lst.append(fn2)
    lst.append(fn3)
    dct[i] = f3

fn1 = f"""
std::string {generate_alphanumeric_string(30)}() {{
"""

flag_var = generate_alphanumeric_string(30)

xns = f"""
std::string {flag_var};
"""

for i in dct.keys():
    x = generate_alphanumeric_string(30)
    fx = "char " + x + " = (char)" + dct[i] + "();"
    fn1 += fx + "\n"
    xns += flag_var + " += " + x + ";\n"

fn1 += f"""
 {xns}
 return {flag_var};
}}
"""

# Randomize the order of the functions
random.shuffle(lst)
random.shuffle(allfn)

print("\n".join(allfn))
print("\n".join(lst))

print(fn1)