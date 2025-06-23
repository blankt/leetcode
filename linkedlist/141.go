package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}

		slow = slow.Next

		if slow == fast {
			return true
		}
	}
	return false
}
