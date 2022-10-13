# What is this example

This is an example of including util.c in each of the a and b packages.
Error due to double definition of symbol (Say) in util.c

# BUILD

```console
go build .
```

## Output result is the following.

```console
# example
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
duplicate symbol '_Say' in:
    /var/folders/09/5r67q5px4s77k_gpftygn1280000gp/T/go-link-1480914570/000013.o
    /var/folders/09/5r67q5px4s77k_gpftygn1280000gp/T/go-link-1480914570/000016.o
ld: 1 duplicate symbol for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```
