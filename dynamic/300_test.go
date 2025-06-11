package dynamic

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		s    string
		nums []int
		want int
	}{
		{"4", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, // Example from the problem statement
		{"4", []int{0, 1, 0, 3, 2, 3}, 4},           // Another example
		{"1", []int{7, 7, 7, 7, 7, 7}, 1},           // All elements the same
		{"6", []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6}, // Increasing and decreasing sequence
		{"1", []int{3, 2}, 1},                       // Decreasing sequence
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			if got != tt.want {
				t.Errorf("lengthOfLIS(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
