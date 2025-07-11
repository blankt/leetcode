package leetcode

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = &ListNode{Val: 5}
	result := removeNthFromEnd(list, 2)
	expected := &ListNode{Val: 1}
	expected.Next = &ListNode{Val: 2}
	expected.Next.Next = &ListNode{Val: 3}
	expected.Next.Next.Next = &ListNode{Val: 5}
	for result != nil && expected != nil {
		if result.Val != expected.Val {
			t.Errorf("Expected %d, got %d", expected.Val, result.Val)
		}
		result = result.Next
		expected = expected.Next
	}
}
