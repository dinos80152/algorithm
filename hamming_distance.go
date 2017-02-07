package main

import (
	"strconv"
	"strings"
)

func HammingDistance(x int, y int) int {
	xorInt := x ^ y
	xorString := strconv.FormatInt(int64(xorInt), 2)
	return strings.Count(xorString, "1")
}
