// TODO: think solution
package main

import "fmt"

func main() {
	fmt.Println(minimumCost([]int{1, 2, 4, 5, 7, 29, 30}))
}

func minimumCost(days []int) int {
	travelDays := make(map[int]bool)
	for _, day := range days {
		travelDays[day] = true
	}

	var dailyAccumulatedCost [31]int
	for i := 1; i <= 30; i++ {
		if !travelDays[i] {
			dailyAccumulatedCost[i] = dailyAccumulatedCost[i-1]
			continue
		}

		// 1 day
		minCost := dailyAccumulatedCost[i-1] + 2

		// 7 days
		for j := max(0, i-7); j <= i-4; j++ {
			minCost = min(minCost, dailyAccumulatedCost[j]+7)
		}

		// 30 days
		minCost = min(minCost, 25)

		dailyAccumulatedCost[i] = minCost
	}
	return dailyAccumulatedCost[30]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
