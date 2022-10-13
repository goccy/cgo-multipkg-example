#include <stdio.h>

extern void Say();

void Say(int v)
{
  fprintf(stderr, "%d\n", v);
}
