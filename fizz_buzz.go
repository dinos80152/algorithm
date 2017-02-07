package main

import "strconv"

func FizzBuzz(n int) []string {

	s := make([]string, n)
	var result string
	for n > 0 {
		index := n - 1
		if n%3 == 0 && n%5 == 0 {
			result = "FizzBuzz"
		} else if n%3 == 0 {
			result = "Fizz"
		} else if n%5 == 0 {
			result = "Buzz"
		} else {
			result = strconv.Itoa(n)
		}
		s[index] = result
		n--
	}
	return s
}
