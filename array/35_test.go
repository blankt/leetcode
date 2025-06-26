package array

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		insert int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{}, 1, 0},
	}
	for _, test := range tests {
		result := searchInsert(test.nums, test.target)
		if result != test.insert {
			t.Errorf("searchInsert(%v, %d) = %d; want %d", test.nums, test.target, result, test.insert)
		}
	}
}
