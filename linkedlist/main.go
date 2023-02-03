package main

//
//import (
//	"container/heap"
//	"fmt"
//)
//
//func main() {
//	nodes := []*ListNode{
//		{Val: 6, Next: nil},
//		{Val: 4, Next: nil},
//		{Val: 2, Next: nil},
//		{Val: 3, Next: nil},
//	}
//
//	pq := PriorityQueue{}
//	for _, v := range nodes {
//		item := &Item{
//			Val:      v,
//			Priority: v.Val,
//		}
//		pq.Push(item)
//	}
//
//	heap.Init(&pq)
//
//	for pq.Len() > 0 {
//		x := pq.Pop()
//		fmt.Println(x)
//	}
//}
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
////func mergeKLists(lists []*ListNode) *ListNode {
////
////}
//
//type Item struct {
//	Val      interface{}
//	Priority int
//	Index    int
//}
//
//type PriorityQueue []*Item
//
//func (pq PriorityQueue) Len() int {
//	return len(pq)
//}
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].Priority < pq[j].Priority
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].Index = i
//	pq[j].Index = j
//}
//
//func (pq *PriorityQueue) Push(x any) {
//	n := len(*pq)
//	item := x.(*Item)
//	item.Index = n
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() any {
//	old := *pq
//	item := old[0]
//	item.Index = -1
//	old[0] = old[old.Len()-1]
//	*pq = old[:old.Len()-1]
//	heap.Fix(pq, 0)
//	return item
//}
