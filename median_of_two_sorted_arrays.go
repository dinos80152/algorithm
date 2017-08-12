package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsL, numsS := nums1, nums2
	if len(nums2) > len(nums1) {
		numsL, numsS = nums2, nums1
	}
	if len(numsS) == 0 {
		return calMedian(numsL, len(numsL))
	}
	nums := []int{}
	start := 0
	end := -1
	for i := 0; i < len(numsS); i++ {
		for j := 0; j < len(numsL); j++ {
			if numsS[i] <= numsL[j] {
				nums = append(nums, numsL[start:j]...)
				nums = append(nums, numsS[i])
				start = j
				end = i
				break
			}
		}
	}
	nums = append(nums, numsL[start:]...)
	if end == -1 || end != len(numsS)-1 {
		nums = append(nums, numsS[end+1:]...)
	}
	return calMedian(nums, len(nums))
}

func calMedian(nums []int, len int) float64 {
	p1 := len / 2
	if len%2 == 0 {
		p2 := p1 - 1
		return (float64(nums[p1]) + float64(nums[p2])) / 2
	}
	return float64(nums[p1])
}
