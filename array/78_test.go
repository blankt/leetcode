package array

import "testing"

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	if len(res) != 8 {
		t.Errorf("expected 8, got %d", len(res))
	}

	res = subsets2(nums)
	if len(res) != 8 {
		t.Errorf("expected 8, got %d", len(res))
	}
}
