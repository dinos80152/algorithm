package main

import (
	"strconv"
	"strings"
)

func HammingDistance(x int, y int) int {
	xBinaryString, yBinaryString := strconv.FormatInt(int64(x), 2), strconv.FormatInt(int64(y), 2)
	xBoolArray, yBoolArray := binaryString2BoolArray(xBinaryString), binaryString2BoolArray(yBinaryString)

	var diff int
	var longSlice, shortSlice []bool

	longSlice, shortSlice = sortByArrayLength(xBoolArray, yBoolArray)

	for index, value := range longSlice {
		if value != shortSlice[index] {
			diff++
		}
	}

	return diff
}

func binaryString2BoolArray(binS string) []bool {
	stringArray := strings.Split(binS, "")
	var boolArray []bool
	boolArray = make([]bool, len(stringArray))
	arrLength := len(stringArray)
	for index, value := range stringArray {
		boolArray[arrLength-index-1], _ = strconv.ParseBool(value)
	}

	return boolArray
}

func sortByArrayLength(xBoolArray []bool, yBoolArray []bool) (longSlice []bool, shortSlice []bool) {
	if len(xBoolArray) >= len(yBoolArray) {
		longSlice = make([]bool, len(xBoolArray))
		shortSlice = make([]bool, len(xBoolArray))
		copy(longSlice, xBoolArray)
		copy(shortSlice, yBoolArray)
	} else {
		longSlice = make([]bool, len(yBoolArray))
		shortSlice = make([]bool, len(yBoolArray))
		copy(longSlice, yBoolArray)
		copy(shortSlice, xBoolArray)
	}
	return longSlice, shortSlice
}
