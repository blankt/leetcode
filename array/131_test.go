package array

import (
	"testing"
)

func TestPartition(t *testing.T) {
	s := "cbbbcc"
	result := partition(s)
	if len(result) != 9 {
		t.Errorf("Expected %d partitions, got %d", 9, len(result))
		return
	}
}
