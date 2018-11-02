package main

import "testing"

func TestFibonacci(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"5", args{5}, 5},
		{"6", args{6}, 8},
		{"7", args{7}, 13},
		{"8", args{8}, 21},
		{"9", args{9}, 34},
		{"10", args{10}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciWithMap(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"5", args{5}, 5},
		{"6", args{6}, 8},
		{"7", args{7}, 13},
		{"8", args{8}, 21},
		{"9", args{9}, 34},
		{"10", args{10}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciWithMap(tt.args.n); got != tt.want {
				t.Errorf("FibonacciWithMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
