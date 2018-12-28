// The Hamming distance between two integers is the number of positions at which
//  the corresponding bits are different.

// Given two integers x and y, calculate the Hamming distance.

package bit

func HammingDistance(x int, y int) int {
	xorInt := x ^ y
	var diff int
	for xorInt > 0 {
		diff++
		xorInt = xorInt & (xorInt - 1)
	}
	return diff
}
