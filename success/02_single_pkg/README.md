# What is this example

This example loads all the source code necessary to use the apis of the library in one package.
If there are common symbols in each source code, rename the symbol name before loading to avoid conflicts.
But, reading all the source code in one package will compile all the source code serially, so it will take a long time to compile...

# BUILD

```console
go build .
```
