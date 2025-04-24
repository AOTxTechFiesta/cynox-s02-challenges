#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

void plumber_warp_gate()
{
  printf("\n🚀 Portal locked in!\n");
  printf("🌍 Welcome back to Earth, Grandpa Max.\n");
  system("cat /flag.txt");
}

int main()
{
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stdin, NULL, _IONBF, 0);

  fflush(stdout);
  fflush(stdin);

  unsigned long null_void_coords;
  char omnitrix_buffer[1024] = {'\0'};

  printf("=== NULL VOID ESCAPE PROTOCOL ===\n\n");
  printf("Enter your name to calibrate your identity: ");

  fgets(omnitrix_buffer, 1024, stdin);
  printf("\nCalibration Accepted: ");
  printf(omnitrix_buffer, &main);
  printf("\n");

  printf("Now... input the coordinates of the return portal: ");
  scanf("%lx", &null_void_coords);
  void (*warp_function)(void) = (void(*)())null_void_coords;

  printf("\nSynchronizing with portal at %lx...\n", null_void_coords);
  printf("Engaging warp sequence...\n");
  
  warp_function();

  return 0;
}