# Yal

Yet another Lisp [transpiler](https://en.wikipedia.org/wiki/Source-to-source_compiler) to Go source code.

[![GoDoc](https://godoc.org/github.com/droptheplot/yal?status.svg)](https://godoc.org/github.com/droptheplot/yal)
[![Go Report Card](https://goreportcard.com/badge/github.com/droptheplot/yal)](https://goreportcard.com/report/github.com/droptheplot/yal)

## Usage

```shell
git clone https://github.com/droptheplot/yal
cd yal
go build
./yal -path hello.yal
```

## Example

**input.yal**

```go
(package "main")

(import "fmt")

(func main () ()
  (fmt.Println (mul 2 3)))

(func mul ((a int) (b int)) (int)
  (* a b))
```

**output.go**

```go
package main

import "fmt"

func main() {
  fmt.Println(mul(2, 3))
}

func mul(a int, b int) int {
  return a * b
}
```

## Disclaimer

This is a fun project and is not meant to be used anywhere.

## Features

- Declarations (`func`, `var`, `package`, `import`)
- Arithmetics (`+`, `-`, `*`, `/`, `%`)
- Comparisons (`<`, `>`, `<=`, `>=`, `==`, `!=`)
- Boolean operators (`&&`, `||`)
- Slice operations (`map`)
- Control flow statements (`if`, `switch`, `return`)
