package bleedingink

import (
	"reflect"
	"testing"
)

func Test_bleedInk(t *testing.T) {

	getCleanSheet := func(h, w int) (sheet [][]int) {
		sheet = make([][]int, h)
		for i := range sheet {
			sheet[i] = make([]int, w)
		}
		return
	}

	type args struct {
		h, w int
		dots []dot
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
		wantTotal  int
	}{
		{"3x3, 1 dots", args{3, 3, []dot{{1, 1, 10}}}, [][]int{
			{8, 9, 8},
			{9, 10, 9},
			{8, 9, 8},
		}, 78},
		{"4x3, 2 dots", args{3, 4, []dot{
			{0, 0, 255},
			{x: 2, y: 1, darkness: 255},
		}}, [][]int{
			{255, 254, 254, 253},
			{254, 254, 255, 254},
			{253, 253, 254, 253},
		}, 3046},
		{"6x5, 4 dots", args{5, 6, []dot{
			{y: 1, x: 0, darkness: 10},
			{y: 2, x: 2, darkness: 9},
			{y: 2, x: 3, darkness: 5},
			{y: 4, x: 2, darkness: 9},
		}}, [][]int{
			{9, 8, 7, 6, 5, 4},
			{10, 9, 8, 7, 6, 5},
			{9, 8, 9, 8, 7, 6},
			{8, 7, 8, 7, 6, 5},
			{7, 8, 9, 8, 7, 6},
		}, 217},
		{"500x500, 1 dots", args{500, 500, []dot{
			{0, 0, 4000},
		}}, [][]int{}, 875250000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotTotal := bleedInk(getCleanSheet(tt.args.h, tt.args.w), tt.args.dots)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				// t.Errorf("totalDarkness() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("totalDarkness() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
