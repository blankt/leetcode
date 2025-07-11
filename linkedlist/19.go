package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}

	dummy := &ListNode{Next: head}

	front := dummy
	back := dummy
	for i := 0; i < n; i++ {
		if back == nil {
			return nil
		}
		back = back.Next
	}
	for back.Next != nil {
		front = front.Next
		back = back.Next
	}
	front.Next = front.Next.Next
	return dummy.Next
}
