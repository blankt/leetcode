package leetcode

import "testing"

func TestSwapPairs(t *testing.T) {
	// Test cases for swapPairs function
	testCases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			input:    nil,
			expected: nil,
		},
		{
			input:    &ListNode{Val: 1},
			expected: &ListNode{Val: 1},
		},
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
		},
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}},
		},
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
		},
	}

	for _, tc := range testCases {
		result := swapPairs(tc.input)
		if !compareLists(result, tc.expected) {
			panic("Test failed")
		}
	}
}

func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
