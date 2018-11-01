package enumeration

import (
	"reflect"
	"testing"
)

func Test_dollarBuyChicken(t *testing.T) {
	type args struct {
		dollar   int
		quantity int
	}
	tests := []struct {
		name string
		args args
		want [][3]int
	}{
		{"100", args{
			100, 100,
		}, [][3]int{
			[3]int{0, 25, 75},
			[3]int{4, 18, 78},
			[3]int{8, 11, 81},
			[3]int{12, 4, 84},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dollarBuyChicken(tt.args.dollar, tt.args.quantity)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dollarBuyChicken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
