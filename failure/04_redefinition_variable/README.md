# What is this example

This avoids the `01` pattern by moving the library source code to a specific location ( `lib` directory ) and referencing only the necessary files.
But We get an error because a.c and b.c use a common variable name.

# BUILD

```console
go build .
```

## Output result is the following.

```console
# example
In file included from bind.c:2:
./lib/b.c:6:12: error: redefinition of 'v'
./lib/a.c:6:12: note: previous definition is here
```
