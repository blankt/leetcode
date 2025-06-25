package leetcode

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	headA := &ListNode{Val: 1}
	headA.Next = &ListNode{Val: 2}
	headA.Next.Next = &ListNode{Val: 3}

	headB := &ListNode{Val: 4}
	headB.Next = headA.Next // Intersection at node with value 2

	intersection := getIntersectionNode(headA, headB)
	if intersection == nil || intersection.Val != 2 {
		t.Errorf("Expected intersection at node with value 2, got %v", intersection)
	}

	headC := &ListNode{Val: 5} // No intersection
	intersection = getIntersectionNode(headA, headC)
	if intersection != nil {
		t.Errorf("Expected no intersection, got %v", intersection)
	}

	headD := &ListNode{Val: 4}
	headD.Next = &ListNode{Val: 1}
	headD.Next.Next = &ListNode{Val: 8}
	headD.Next.Next.Next = &ListNode{Val: 4}
	headD.Next.Next.Next.Next = &ListNode{Val: 5}
	headE := &ListNode{Val: 5}
	headE.Next = &ListNode{Val: 6}
	headE.Next.Next = &ListNode{Val: 1}
	headE.Next.Next.Next = headD.Next.Next
	intersection = getIntersectionNode(headD, headE)
	if intersection == nil || intersection.Val != 8 {
		t.Errorf("Expected intersection at node with value 8, got %v", intersection)
	}
}
