package leetcode

import "testing"

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	result := mergeTwoLists(list1, list2)

	length := 0
	for result != nil {
		length++
		result = result.Next
	}

	if length != 6 {
		t.Fatalf("Expected length 7, got %d", length)
	}
}
