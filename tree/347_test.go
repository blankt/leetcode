package tree

import "testing"

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}

	for _, test := range tests {
		got := topKFrequent(test.nums, test.k)
		if len(got) != len(test.want) {
			t.Errorf("topKFrequent(%v, %d) = %v; want %v", test.nums, test.k, got, test.want)
			continue
		}
		for i, v := range got {
			if v != test.want[i] {
				t.Errorf("topKFrequent(%v, %d) = %v; want %v", test.nums, test.k, got, test.want)
				break
			}
		}
	}
}
