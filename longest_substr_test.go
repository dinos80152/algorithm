package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"c", 1},
		{"dvdf", 3},
	}
	for _, tt := range tests {
		if got := lengthOfLongestSubstring(tt.args); got != tt.want {
			t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
		}
	}
}
