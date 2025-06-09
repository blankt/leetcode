package dynamic

import "testing"

func TestNumSquares(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{n: 12, expected: 3}, // 4 + 4 + 4
		{n: 13, expected: 2}, // 9 + 4
		{n: 1, expected: 1},  // 1
		{n: 2, expected: 2},  // 1 + 1
	}

	for _, test := range tests {
		result := numSquares(test.n)
		if result != test.expected {
			t.Errorf("numSquares(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}
