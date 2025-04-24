import base64

flag = "cyn0x{n0t_s3cur3_just_3nc0d3d_l1k3_a_r0gu3_galvan_pr0t0c0l}"

encoded_flag = base64.b64encode(flag.encode()).decode()

open("flag.txt", "w").write(encoded_flag)