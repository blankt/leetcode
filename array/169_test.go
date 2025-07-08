package array

import "testing"

func TestMajorityNums(t *testing.T) {
	nums := []int{3, 2, 3}
	if majorityElement(nums) != 3 {
		t.Errorf("Expected 3, got %d", majorityElement(nums))
	}

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	if majorityElement(nums) != 2 {
		t.Errorf("Expected 2, got %d", majorityElement(nums))
	}

	nums = []int{6, 6, 6, 7, 7}
	if majorityElement(nums) != 6 {
		t.Errorf("Expected 0 for empty array, got %d", majorityElement(nums))
	}
}
