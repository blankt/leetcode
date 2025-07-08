package leetcode

import "testing"

func TestCopyRandomList(t *testing.T) {
	head := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
			},
		},
		Random: nil,
	}

	head.Next.Random = head
	head.Next.Next.Random = head.Next

	copiedList := copyRandomList(head)

	if copiedList.Val != 1 || copiedList.Next.Val != 2 || copiedList.Next.Next.Val != 3 {
		t.Errorf("Copied list values do not match original")
	}

	if copiedList.Next.Random != copiedList || copiedList.Next.Next.Random != copiedList.Next {
		t.Errorf("Random pointers do not match original")
	}
}
