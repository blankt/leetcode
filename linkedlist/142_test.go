package leetcode

import "testing"

func TestDetectCycle(t *testing.T) {
	var tests []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}

	tests = append(tests, struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		name:     "head is nil",
		head:     nil,
		expected: nil,
	})

	testNode := &ListNode{Val: 1}
	testNode.Next = &ListNode{Val: 2}
	testNode.Next.Next = &ListNode{Val: 3}
	testNode.Next.Next.Next = &ListNode{Val: 4}
	testNode.Next.Next.Next.Next = testNode.Next
	tests = append(tests, struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		name:     "cycle exists",
		head:     testNode,
		expected: testNode.Next,
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := detectCycle(test.head)
			if result != test.expected {
				t.Errorf("For input %v expected %v but got %v", test.head, test.expected, result)
			}
		})
	}
}
