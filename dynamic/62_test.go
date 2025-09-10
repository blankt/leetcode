package dynamic

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m        int
		n        int
		expected int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
	}

	for _, test := range tests {
		result := uniquePaths1(test.m, test.n)
		if result != test.expected {
			t.Errorf("uniquePaths(%d, %d) = %d; expected %d", test.m, test.n, result, test.expected)
		}
	}
}
