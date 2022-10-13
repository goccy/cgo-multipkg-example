#include <stdio.h>

extern void FuncA();
void Say(int v);

static int v = 1;

void FuncA()
{
  Say(v);
}
