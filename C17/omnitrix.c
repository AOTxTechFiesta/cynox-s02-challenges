#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void transform_into_xlr8()
{
  printf("Accessing Omnitrix Core...\n");
  FILE *fptr;
  char c;

  fptr = fopen("/flag.txt", "r");
  if (fptr == NULL)
  {
    printf("Cannot open file.\n");
    exit(0);
  }

  c = fgetc(fptr);
  while (c != EOF)
  {
    printf("%c", c);
    c = fgetc(fptr);
  }

  printf("\n");
  fclose(fptr);
}

void interface()
{
  char name[64];
  printf("State your name, wielder of the Omnitrix: ");
  fgets(name, 1024, stdin);
  printf("Welcome, %s\n", name);
}

int main()
{
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stdin, NULL, _IONBF, 0);

  fflush(stdout);
  fflush(stdin);
  
  printf("=== Omnitrix Authorization Interface ===\n");
  interface();
  printf("Authorization Failed. Alien transformation denied.\n");
  return 0;
}
