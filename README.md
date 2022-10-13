# cgo-multipkg-example

This contains of issues and solutions for binding multi-package libraries with cgo.
Knowledge acquired in creating [go-graphviz](https://github.com/goccy/go-graphviz) and [go-zetasql](https://github.com/goccy/go-zetasql)

### File Layout

├── LICENSE
├── README.md
├── failure
│   ├── 01_unnecessary_source
│   ├── 02_duplicated_symbol
│   ├── 03_undefined_symbol
│   └── 04_redefinition_variable
├── lib
│   ├── Makefile
│   ├── a.c
│   ├── b.c
│   ├── test.c
│   └── util.c
└── success
    ├── 01_replace_symbol
    ├── 02_single_pkg
    └── 03_export_symbol
