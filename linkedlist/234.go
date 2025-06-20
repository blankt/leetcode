package leetcode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	temp = head
	for i := 0; i < length/2-1; i++ {
		temp = temp.Next
	}

	middle := temp
	if length%2 == 1 {
		middle = middle.Next.Next
	} else {
		middle = middle.Next
	}
	temp.Next = nil

	reverse := reverseList(middle)
	for reverse != nil && head != nil {
		if reverse.Val != head.Val {
			return false
		}
		reverse = reverse.Next
		head = head.Next
	}
	if reverse != nil || head != nil {
		return false
	}

	return true
}
