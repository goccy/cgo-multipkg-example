# What is this example

This example refer to util.c from package a but not from b.
As a result, processing by the linker of package b causes an undefined symbol error.

# BUILD

```console
go build .
```

## Output result is the following.

```console
# example/b
Undefined symbols for architecture x86_64:
  "_Say", referenced from:
      _FuncB in _x003.o
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```
