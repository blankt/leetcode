package point

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}

	for _, test := range tests {
		result := maxArea(test.height)
		if result != test.expected {
			t.Errorf("For input %v expected %d but got %d", test.height, test.expected, result)
		}
	}
}
