package array

import "testing"

func Benchmark_TwoSum(b *testing.B) {
	arr := []int{2, 7, 11, 15}
	t := 22
	for i := 0; i < b.N; i++ {
		TwoSum(arr, t)
	}
}

func Benchmark_TwoSumWithMap(b *testing.B) {
	arr := []int{2, 7, 11, 15}
	t := 22
	for i := 0; i < b.N; i++ {
		TwoSumWithMap(arr, t)
	}
}
