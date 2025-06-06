package dynamic

import "testing"

func TestRob(t *testing.T) {
	nums := []int{0, 3, 2, 0}
	expected := 3

	result := rob(nums)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
