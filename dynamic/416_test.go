package dynamic

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{2, 2, 3, 5}, false},
	}

	for _, test := range tests {
		result := canPartition(test.nums)
		if result != test.expected {
			t.Errorf("canPartition(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}
