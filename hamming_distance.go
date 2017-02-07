package main

func HammingDistance(x int, y int) int {
	xorInt := x ^ y
	var diff int
	for xorInt > 0 {
		diff++
		xorInt = xorInt & (xorInt - 1)
	}
	return diff
}
