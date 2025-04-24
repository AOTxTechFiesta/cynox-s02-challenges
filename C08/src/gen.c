#include <stdio.h>
#include <string.h>

int main() {
  unsigned char ENCRYPTED_FLAG[] = {
    0x21, 0x3b, 0x2c, 0x72, 0x3a, 0x39, 0x34, 0x73, 0x2e, 0x25, 0x23, 0x3a, 0x1d, 0x71, 0x2c, 0x21, 0x72, 0x26, 0x71, 0x26, 0x1d, 0x3b, 0x72, 0x37, 0x1d, 0x2f, 0x37, 0x31, 0x36, 0x1d, 0x30, 0x71, 0x34, 0x71, 0x30, 0x31, 0x71, 0x1d, 0x73, 0x36, 0x3f
  };

  const int FLAG_LEN = (sizeof(ENCRYPTED_FLAG) / sizeof(char));
  const char XOR_KEY = 0x42;

  char input[100];
  char decrypted_flag[(sizeof(ENCRYPTED_FLAG) / sizeof(char)) + 1];

  for (int i = 0; i < FLAG_LEN; i++) {
    decrypted_flag[i] = ENCRYPTED_FLAG[i] ^ XOR_KEY;
  }
  decrypted_flag[FLAG_LEN] = '\0';

  printf("Enter the password: ");
  if (fgets(input, sizeof(input), stdin) != NULL) {
    size_t len = strlen(input);
    if (len > 0 && input[len - 1] == '\n')
      input[len - 1] = '\0';

    if (strcmp(input, decrypted_flag) == 0) {
      printf("Hooray! You found the flag! Now submit it to get your points before anyone does!\n");
    } else {
      printf("Nope! That's not it.\n");
    }
  }

  return 0;
}
