package testdata

import "strconv"

func FizzBuzz(numbers []int) []string {
	return func(values []int) []string {
		newValues := make([]string, len(values))
		for i := range values {
			newValues[i] = toFizzBuzz(values[i])
		}
		return newValues
	}(numbers)
}
func toFizzBuzz(n int) string {
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
