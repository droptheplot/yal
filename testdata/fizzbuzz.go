package main

import (
	"fmt"
	"strconv"
)

func main() {
	FizzBuzz(1, 15)
}
func FizzBuzz(n int, e int) {
	if n == e+1 {
		return
	}
	fmt.Println(stringer(n))
	FizzBuzz(n+1, e)
}
func stringer(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return ("FizzBuzz")
	}
	if n%3 == 0 {
		return ("Fizz")
	}
	if n%5 == 0 {
		return ("Buzz")
	}
	return (strconv.Itoa(n))
}
