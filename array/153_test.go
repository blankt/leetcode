package array

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{2, 1}, expected: 1},
		{nums: []int{1}, expected: 1},
		{nums: []int{3, 4, 5, 1, 2}, expected: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, expected: 0},
		{nums: []int{11, 13, 15, 17}, expected: 11},
	}

	for _, test := range tests {
		result := findMin1(test.nums)
		if test.expected != result {
			t.Errorf("expected: %d, got: %d", test.expected, result)
		}
	}
}
