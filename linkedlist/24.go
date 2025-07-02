package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	dummyHead := dummy
	first, second := head, head.Next
	for second != nil {
		temp := second.Next
		second.Next = first
		first.Next = temp
		dummyHead.Next = second
		dummyHead = first

		if temp == nil || temp.Next == nil {
			break
		}
		first = temp
		second = temp.Next
	}

	return dummy.Next
}
