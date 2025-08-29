package tree

import "container/heap"

type FrequentK struct {
	Times int
	Num   int
}

type FrequentKHeap struct {
	data []FrequentK
	size int
}

func (h *FrequentKHeap) Less(i, j int) bool {
	return (*h).data[i].Times < (*h).data[j].Times
}

func (h *FrequentKHeap) Swap(i, j int) {
	(*h).data[i], (*h).data[j] = (*h).data[j], (*h).data[i]
}

func (h *FrequentKHeap) Len() int {
	return len((*h).data)
}

func (h *FrequentKHeap) Push(v any) {
	(*h).data = append((*h).data, v.(FrequentK))
}

func (h *FrequentKHeap) Pop() (v any) {
	(*h).data, v = (*h).data[:h.Len()-1], (*h).data[h.Len()-1]
	return
}

func topKFrequent(nums []int, k int) []int {
	items := make(map[int]FrequentK)
	for _, num := range nums {
		if item, ok := items[num]; ok {
			item.Times++
			items[num] = item
		} else {
			items[num] = FrequentK{
				Times: 1,
				Num:   num,
			}
		}
	}

	myHeap := &FrequentKHeap{
		size: k,
	}
	for _, v := range items {
		item := FrequentK{
			Times: v.Times,
			Num:   v.Num,
		}
		heap.Push(myHeap, item)
		if myHeap.Len() > k {
			heap.Pop(myHeap)
		}
	}
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		v := heap.Pop(myHeap)
		item, _ := v.(FrequentK)
		result = append(result, item.Num)
	}
	return result
}
