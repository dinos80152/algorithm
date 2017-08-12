package main

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"", args{[]int{}, []int{2}}, 2},
		{"", args{[]int{2}, []int{}}, 2},
		{"", args{[]int{1}, []int{1}}, 1},
		{"", args{[]int{10001}, []int{10000}}, 10000.5},
		{"", args{[]int{1, 3}, []int{2}}, 2},
		{"", args{[]int{1, 3}, []int{2, 4}}, 2.5},
		{"", args{[]int{1, 2}, []int{1, 1}}, 1},
		{"", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		{"", args{[]int{3, 4}, []int{1, 2}}, 2.5},
		{"", args{[]int{1, 2, 8, 10, 98}, []int{3, 4, 9, 10, 25, 44}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
