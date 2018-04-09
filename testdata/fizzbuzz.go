package testdata

import "strconv"

func FizzBuzz(n int, end int, res *[]string) {
	if n == end+1 {
		return
	}
	*res = append(*res, stringer(n))
	FizzBuzz(n+1, end, res)
}
func stringer(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
