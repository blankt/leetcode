package greedy

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{1, 0, 0}, false},
		{[]int{2, 0, 0}, true},
		{[]int{1, 2, 3, 4, 5}, true},
	}

	for _, test := range tests {
		result := canJump(test.nums)
		if result != test.expected {
			t.Errorf("canJump(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}
