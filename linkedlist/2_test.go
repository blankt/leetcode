package leetcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{}
	l1.Val = 2
	l1.Next = &ListNode{Val: 4, Next: &ListNode{Val: 3}}

	l2 := &ListNode{}
	l2.Val = 5
	l2.Next = &ListNode{Val: 6, Next: &ListNode{Val: 4}}

	result := addTwoNumbers(l1, l2)

	expected := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}

	for result != nil && expected != nil {
		if result.Val != expected.Val {
			t.Errorf("Expected %d, got %d", expected.Val, result.Val)
		}
		result = result.Next
		expected = expected.Next
	}

	if result != nil || expected != nil {
		t.Error("Lists are of different lengths")
	}
}
