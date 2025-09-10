package dynamic

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}

	for _, test := range tests {
		result := minPathSum(test.grid)
		if result != test.expected {
			t.Errorf("minPathSum(%+v) = %d, want %d", test.grid, result, test.expected)
		}
	}
}
