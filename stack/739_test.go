package stack

import (
	"reflect"
	"testing"
)

func TestDailyTemperature(t *testing.T) {
	tests := []struct {
		temperature []int
		want        []int
	}{
		{
			[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			temperature: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:        []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperature: []int{30, 40, 50, 60},
			want:        []int{1, 1, 1, 0},
		},
		{
			temperature: []int{30, 60, 90},
			want:        []int{1, 1, 0},
		},
	}

	for _, test := range tests {
		res := dailyTemperatures2(test.temperature)
		if !reflect.DeepEqual(res, test.want) {
			t.Errorf("dailyTemperatures => %v, want %v", res, test.want)
		}
	}
}
