package array

import (
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	if len(res) != 6 {
		t.Errorf("expected 6, got %d", len(res))
	}
}
