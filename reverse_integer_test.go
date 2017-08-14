package main

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{0, 0},
		{123, 321},
		{1000000, 1},
		{1000020, 200001},
		{-321, -123},
		{1000000003, 0},
	}
	for _, tt := range tests {
		if got := reverse(tt.arg); got != tt.want {
			t.Errorf("reverse() = %v, want %v", got, tt.want)
		}
	}
}
