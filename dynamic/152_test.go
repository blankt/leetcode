package dynamic

import "testing"

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	expected := 6
	result := maxProduct(nums)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

	nums = []int{-2, 0, -1}
	expected = 0
	result = maxProduct(nums)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

	nums = []int{-2, 3, -4}
	expected = 24
	result = maxProduct(nums)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
