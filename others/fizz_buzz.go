// Write a program that outputs the string representation of numbers from 1 to n.

// But for multiples of three it should output “Fizz” instead of the number and
// for the multiples of five output “Buzz”. For numbers which are multiples of
// both three and five output “FizzBuzz”.

// https://leetcode.com/problems/fizz-buzz/

package others

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
