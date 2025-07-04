package leetcode

import "testing"

func TestSortList(t *testing.T) {
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
			input:    &ListNode{Val: 2, Next: &ListNode{Val: 1}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			input:    &ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			expected: &ListNode{Val: -1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}},
		},
	}

	for _, tc := range testCases {
		result := sortList(tc.input)
		if !compareLists(result, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}
