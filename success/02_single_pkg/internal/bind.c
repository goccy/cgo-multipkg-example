#define v v_a
#include "a.c"
#undef v

#define v v_b
#include "b.c"
#undef v

#include "util.c"
