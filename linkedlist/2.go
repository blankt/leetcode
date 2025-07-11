package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	current := dummy
	carry := 0

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		current.Next = node
		current = node
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		for l1 != nil {
			sum := l1.Val + carry
			node := &ListNode{Val: sum % 10}
			current.Next = node
			current = node
			carry = sum / 10
			l1 = l1.Next
		}
	}

	if l2 != nil {
		for l2 != nil {
			sum := l2.Val + carry
			node := &ListNode{Val: sum % 10}
			current.Next = node
			current = node
			carry = sum / 10
			l2 = l2.Next
		}
	}

	if carry > 0 {
		node := &ListNode{Val: carry}
		current.Next = node
	}
	return dummy.Next
}
