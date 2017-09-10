# Yal

Yet another Lisp.

[![GoDoc](https://godoc.org/github.com/droptheplot/yal?status.svg)](https://godoc.org/github.com/droptheplot/yal)
[![Go Report Card](https://goreportcard.com/badge/github.com/droptheplot/yal)](https://goreportcard.com/report/github.com/droptheplot/yal)

# Usage

```shell
git clone https://github.com/droptheplot/yal
cd yal
go build
./yal -file=hello.yal
```

# Functions

* [Core](#core)
* [Integer](#integer)

## Core

### `==`

```lisp
(== 2 2)
;; true

(== "hello" "world")
;; false
```

### `if`

```lisp
(if (== 2 4)
    (print "true")
    (print "false"))
;; nil

(if (== 2 2)
    (print "equal"))
;; nil
```

### `print`

```lisp
(print "hello world")
;; nil
```

### `def`

```lisp
(def my_var 5)
;; nil
```

### `repeat`

```lisp
(repeat n 5
  (print n))
;; nil
```

## Integer

### `+`

```lisp
(+ 2 2)
;; 4
```

### `-`

```lisp
(- 10 5)
;; 5
```

### `*`

```lisp
(* 2 2)
;; 4
```

### `/`

```lisp
(/ 10 5)
;; 2
```

### `%`

```lisp
(% 3 2)
;; 1
```

### `>`

```lisp
(> 10 5)
;; true
```

### `>=`

```lisp
(>= 10 5)
;; true
```

### `<`

```lisp
(< 10 5)
;; false
```

### `<=`

```lisp
(<= 10 5)
;; false
```
