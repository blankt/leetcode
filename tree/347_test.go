package tree

import "testing"

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}, 3, []int{10, 7, 3}},
		{[]int{4, 1, -1, 2, -1, 2, 3}, 2, []int{-1, 2}},
		{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{2, 1}},
		{[]int{1}, 1, []int{1}},
	}

	for _, test := range tests {
		got := topKFrequent(test.nums, test.k)
		if len(got) != len(test.want) {
			t.Errorf("topKFrequent(%v, %d) = %v; want %v", test.nums, test.k, got, test.want)
			continue
		}
		for _, v := range got {
			exist := false
			for _, w := range test.want {
				if v == w {
					exist = true
					break
				}
			}
			if !exist {
				t.Errorf("topKFrequent(%v, %d) = %v; want %v", test.nums, test.k, got, test.want)
				break
			}
		}
	}
}
