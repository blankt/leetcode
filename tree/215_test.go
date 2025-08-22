package tree

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20, 2},
		{[]int{99, 99}, 1, 99},
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
	}

	for _, tt := range tests {
		if got := findKthLargest(tt.nums, tt.k); got != tt.want {
			t.Errorf("findKthLargest(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
