// Package main has not finish
// TODO: best solution
package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{2, 3, 0}))
}

func missingNumber(nums []int) int {
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
