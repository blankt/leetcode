package leetcode

func Reverse(a *ListNode) *ListNode {
	var (
		pre, cur, next *ListNode
	)
	pre = nil
	cur = a
	next = a

	for cur != nil {
		next = a.Next
		a.Next = pre
		pre = a
		a = next
	}
	return pre
}

//func reverseKGroup(head *ListNode, k int) *ListNode {
//
//}
