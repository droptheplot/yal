# Yal

Yet another Lisp [transpiler](https://en.wikipedia.org/wiki/Source-to-source_compiler) to Go source code.

[![GoDoc](https://godoc.org/github.com/droptheplot/yal?status.svg)](https://godoc.org/github.com/droptheplot/yal)
[![Go Report Card](https://goreportcard.com/badge/github.com/droptheplot/yal)](https://goreportcard.com/report/github.com/droptheplot/yal)

## Usage

```shell
git clone https://github.com/droptheplot/yal
cd yal
go build
./yal -file hello.yal
```

## Example

### Input

```clojure
(func main () ()
  (fmt.Println (mul 2 3)))

(func mul ((a int) (b int)) (int)
  (return (* a b)))
```

### Output

```go
package main

func main() {
	fmt.Println(mul(2, 3))
}

func mul(a int, b int) int {
	return(a * b)
}
```
