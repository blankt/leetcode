package array

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			[]int{5, 1, 3},
			5,
			0,
		},
		{
			[]int{1, 3},
			1,
			0,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
	}

	for _, test := range tests {
		result := search(test.nums, test.target)
		if result != test.expected {
			t.Errorf("search(%v, %d) = %d; expected %d", test.nums, test.target, result, test.expected)
		}
	}
}
