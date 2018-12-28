package string

import "testing"

func Test_LengthOfLongestSubstring(t *testing.T) {
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
		if got := LengthOfLongestSubstring(tt.args); got != tt.want {
			t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
		}
	}
}
