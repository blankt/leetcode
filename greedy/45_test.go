package greedy

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 1, 1, 1, 1}, 4},
	}

	for _, test := range tests {
		result := jump2(test.nums)
		if result != test.expected {
			t.Errorf("jump(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
