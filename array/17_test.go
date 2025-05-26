package array

import "testing"

func TestLetterCombinations(t *testing.T) {
	digits := "23"
	result := letterCombinations(digits)
	if len(result) != 9 {
		t.Errorf("expected 9 combinations, got %d", len(result))
	}
}
