package main

import "container/heap"

//leetcode submit region begin(Prohibit modification and deletion)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := PriorityQueue{}
	dummy := &ListNode{}
	p := dummy

	for _, head := range lists {
		item := &Item{
			Val:      head,
			Priority: head.Val,
		}
		pq.Push(item)
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := pq.Pop().(*Item).Val.(*ListNode)
		if item.Next != nil {
			newItem := &Item{
				Val:      item.Next,
				Priority: item.Val,
			}
			pq.Push(newItem)
		}
		p.Next = item
		p = p.Next
	}
	return dummy.Next
}

type Item struct {
	Val      interface{}
	Priority int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
	heap.Fix(pq, pq.Len())
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	item := old[0]
	item.Index = -1
	old[0] = old[old.Len()-1]
	*pq = old[:old.Len()-1]
	heap.Fix(pq, 0)
	return item
}

//leetcode submit region end(Prohibit modification and deletion)
