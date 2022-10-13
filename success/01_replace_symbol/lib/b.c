#include <stdio.h>

extern void FuncB();
void Say(int v);

static int v = 2;

void FuncB()
{
  Say(v);
}
