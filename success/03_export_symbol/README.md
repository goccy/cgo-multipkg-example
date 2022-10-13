# What is this example

This example makes it possible to use the functions of util.c used by the a and b packages from Go,
and then exposes those functions to the C side using the export directive,
so that the symbols of util.c can be reused.

# BUILD

```console
go build .
```
