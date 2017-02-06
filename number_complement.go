package main

import (
	"strconv"
	"strings"
)

func FindComplement(num int) int {
	numString := strconv.FormatInt(int64(num), 2)
	numArray := strings.Split(numString, "")
	complementSlice := make([]string, len(numArray))
	for index, value := range numArray {
		value, _ := strconv.ParseInt(value, 0, 0)
		if value == 0 {
			complementSlice[index] = "1"
		} else {
			complementSlice[index] = "0"
		}
	}
	complement, _ := strconv.ParseInt(strings.Join(complementSlice, ""), 2, 0)
	return int(complement)
}
