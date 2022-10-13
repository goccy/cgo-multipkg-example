# What is this example

This is an example when all the contents of the library are moved and do binding.
The library has a file containing the main function for test, so symbols conflict when built.


# BUILD

```console
go build .
```

## Output result is the following.

```console
# example
duplicate symbol '_main' in:
    $WORK/b001/_cgo_main.o
    $WORK/b001/_x005.o
ld: 1 duplicate symbol for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```
