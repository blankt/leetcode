package array

import "testing"

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{nums: []int{0, 1, 2, 3}, expected: []int{6, 0, 0, 0}},
		{nums: []int{1, 1, 1, 1}, expected: []int{1, 1, 1, 1}},
		{nums: []int{1, 2, 3}, expected: []int{6, 3, 2}},
		{nums: []int{5, 6, 7}, expected: []int{42, 35, 30}},
	}

	for _, test := range tests {
		result := productExceptSelf(test.nums)
		if len(result) != len(test.expected) {
			t.Errorf("Expected length %d, got %d", len(test.expected), len(result))
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("productExceptSelf(%v) = %v; want %v", test.nums, result, test.expected)
				break
			}
		}
	}
}
