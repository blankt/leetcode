package leetcode

/*
1->2->3
1 <- 2 <- 3
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	middle := head.Next
	pre.Next = nil
	for middle != nil {
		var temp *ListNode
		if middle.Next != nil {
			temp = middle.Next
		} else {
			temp = nil
		}
		middle.Next = pre
		pre = middle
		middle = temp
	}

	return pre
}
