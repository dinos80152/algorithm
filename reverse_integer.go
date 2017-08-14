package main

import "math"

func reverse(x int) int {
	var ans int
	for {
		ans = ans*10 + x%10
		x = x / 10
		if ans > math.MaxInt32 || ans < math.MinInt32 {
			return 0
		}
		if x == 0 {
			break
		}
	}
	return ans
}
