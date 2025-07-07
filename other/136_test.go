package other

import "testing"

func TestSingleNumber(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	expected := 4
	result := singleNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
