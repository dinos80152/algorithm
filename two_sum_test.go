package main

import "testing"

func Benchmark_twoSum(b *testing.B) {
	arr := []int{2, 7, 11, 15}
	t := 22
	for i := 0; i < b.N; i++ {
		twoSum(arr, t)
	}
}

func Benchmark_twoSumWithMap(b *testing.B) {
	arr := []int{2, 7, 11, 15}
	t := 22
	for i := 0; i < b.N; i++ {
		twoSumWithMap(arr, t)
	}
}
