package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA, lengthB := 0, 0
	for node := headA; node != nil; node = node.Next {
		lengthA++
	}
	for node := headB; node != nil; node = node.Next {
		lengthB++
	}
	if lengthA > lengthB {
		headA = forward(headA, lengthA-lengthB)
	} else if lengthB > lengthA {
		headB = forward(headB, lengthB-lengthA)
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func forward(head *ListNode, steps int) *ListNode {
	for steps > 0 && head != nil {
		head = head.Next
		steps--
	}
	return head
}
