package main

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumWithMap(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if r, ok := m[target-num]; ok {
			return []int{r, i}
		}
		m[num] = i
	}
	return nil
}
