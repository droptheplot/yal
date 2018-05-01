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
  (return (* a b)))
```

**output.go**

```go
package main

import "fmt"

func main() {
	fmt.Println(mul(2, 3))
}

func mul(a int, b int) int {
	return(a * b)
}
```

## Disclaimer

- Since `yal` is not a stand-alone language and just Lisp syntax fo Go we can't get rid (at least for now) from all statements: that means it's not possible to make `if`, `switch` return anything like you would in real functional languages.
- This is fun project and is not meant to be used anywhere.

## Features

- Declarations (`func`, `var`, `package`, `import`)
- Arithmetics (`+`, `-`, `*`, `/`, `%`)
- Comparisons (`<`, `>`, `<=`, `>=`, `==`, `!=`)
- Control flow (`if`, `switch`, `return`)
- Boolean operators (`&&`, `||`)
