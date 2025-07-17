package array

import "testing"

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, []int{0, 9}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{2, 2}},
		{[]int{1, 2, 2, 2, 3}, 2, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := searchRange(tt.nums, tt.target)
			if len(got) != len(tt.want) || got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("searchRange(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
