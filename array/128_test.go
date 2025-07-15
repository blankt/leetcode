package array

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 1, 2, 3, 4, 5}, 6},
		{[]int{1, 2, 0, 1}, 3},
		{[]int{1, 3, 5, 7}, 1},
		{[]int{}, 0},
		{[]int{0, -1}, 2},
		{[]int{-100000000, 1, 2, 3, 4, 999999, 100000000}, 4},
	}

	for _, test := range tests {
		result := longestConsecutive(test.nums)
		if result != test.expected {
			t.Errorf("longestConsecutive(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
