// Given an array of integers, return indices of the two numbers such that they add
//  up to a specific target.

// You may assume that each input would have exactly one solution, and you may not
//  use the same element twice.
//
// https://leetcode.com/problems/two-sum/

package array

func TwoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumWithMap(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if r, ok := m[target-num]; ok {
			return []int{r, i}
		}
		m[num] = i
	}
	return nil
}
