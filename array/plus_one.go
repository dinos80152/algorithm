/*
Given a **non-empty** array of digits representing a non-negative integer, plus
one to the integer.

The digits are stored such that the most significant digit is at the head of the
 list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0
 itself.

https://leetcode.com/problems/plus-one/
*/

package array

// PlusOne is 100% answer on leetcode
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
