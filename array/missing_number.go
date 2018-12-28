// Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find
// the one that is missing from the array.
// https://leetcode.com/problems/missing-number/
// TODO: best solution

package array

func MissingNumber(nums []int) int {
	exists := make([]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		exists[nums[i]] = true
	}
	for i, exist := range exists {
		if !exist {
			return i
		}
	}
	return -1
}
