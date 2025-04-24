from stegano import lsb

# Hide the message
secret = lsb.hide("input.png", "cyn0x{v1lg4x_w4s_n3v3r_ju57_4_f4c3_h3_w45_4_m3554g3}")
secret.save("vilgax.png")

print("Message hidden in vilgax.png")
