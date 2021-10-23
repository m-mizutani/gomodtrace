# gomodtrace

Trace go.mod dependencies.

```sh
$ go mod graph | gomodtrace github.com/kr/text@v0.1.0
github.com/kr/text@v0.1.0
+ github.com/kr/pretty@v0.2.1
  + gopkg.in/check.v1@v1.0.0-20201130134442-10cb98267c6c
    + github.com/m-mizutani/goerr@v0.1.2
      + github.com/m-mizutani/gomodtrace
      + github.com/m-mizutani/zlog@v0.0.0-20211023032251-9e5ad6f1e0dc
        + github.com/m-mizutani/gomodtrace```
```

## Install

```sh
$ go get github.com/m-mizutani/gomodtrace
```

## Usage

```sh
NAME:
   gomodtrace - Trace go module dependency by output of go mod graph

USAGE:
   go mod graph | gomodtrace

VERSION:
   v0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input value, -i value   Input file (use '-' for stdin) (default: "-")
   --format value, -f value  Output format [tree|json] (default: "tree")
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

## License

[MIT License](./LICENSE)