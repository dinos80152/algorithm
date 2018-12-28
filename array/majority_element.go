// TODO: best solution
// Given an array of size n, find the majority element. The majority element is
// the element that appears more than ⌊ n/2 ⌋ times.

// You may assume that the array is non-empty and the majority element always
// exist in the array.

// https://leetcode.com/problems/majority-element/

package array

func MajorityElement(nums []int) int {
	countMap := make(map[int]int)
	var most int
	for _, num := range nums {
		countMap[num]++
		if countMap[num] > countMap[most] {
			most = num
		}
	}
	return most
}
