package main

import "testing"

func Test_strstr(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"start", args{"good evening roy", "good"}, 0},
		{"end", args{"good evening roy", "y"}, 15},
		{"middle", args{"good evening roy", "even"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strstr(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("strstr() = %v, want %v", got, tt.want)
			}
		})
	}
}
