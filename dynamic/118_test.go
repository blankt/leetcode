package dynamic

import "testing"

func TestGenerate(t *testing.T) {
	numRows := 5
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

	result := generate(numRows)
	if len(result) != numRows {
		t.Errorf("Expected %d rows, got %d", numRows, len(result))
		return
	}

	for i := range result {
		if len(result[i]) != i+1 {
			t.Errorf("Row %d expected length %d, got %d", i, i+1, len(result[i]))
			return
		}
		for j := range result[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("Value at row %d column %d expected %d, got %d", i, j, expected[i][j], result[i][j])
			}
		}
	}
}
