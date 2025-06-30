package point

import "testing"

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		moveZeroes(test.nums)
		for i := range test.nums {
			if test.nums[i] != test.expected[i] {
				t.Errorf("For input %v expected %v but got %v", test.nums, test.expected, test.nums)
			}
		}
	}

}
