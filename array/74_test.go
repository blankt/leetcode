package array

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
				{60, 61, 62, 63},
			},
			target:   3,
			expected: true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
				{60, 61, 62, 63},
			},
			target:   13,
			expected: false,
		},
	}

	for _, test := range tests {
		result := searchMatrix(test.matrix, test.target)
		if result != test.expected {
			t.Errorf("searchMatrix(%v, %d) = %v; expected %v", test.matrix, test.target, result, test.expected)
		}
	}
}
