#include <stdio.h>
#include <string.h>

const char *flag = "cyn0x{v1lgax_und3r3stim4t3d_th3_p0w3r_0f_y0ur_t00ls}";

void success()
{
  printf("Hooray! You found the flag! Now submit it to get your points before anyone does!\n");
}

int main()
{
  char input[100];

  printf("Enter the password: ");

  if (fgets(input, sizeof(input), stdin) != NULL)
  {
    size_t len = strlen(input);
    if (len > 0 && input[len - 1] == '\n')
    {
      input[len - 1] = '\0';
    }

    if (strcmp(input, flag) == 0)
    {
      success();
    }
    else
    {
      printf("Nope! That's not it.\n");
    }
  }
  else
  {
    printf("Input error!\n");
  }

  return 0;
}
