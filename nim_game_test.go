package main

import "testing"

func TestCanWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test", args{4}, false,
		},
		{
			"test2", args{11}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanWinNim(tt.args.n); got != tt.want {
				t.Errorf("CanWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
