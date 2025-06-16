package greedy

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test:%v", i), func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
			}
		})
	}
}
