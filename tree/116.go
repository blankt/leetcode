package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	leftNode := root
	queue := make([]*Node, 1<<12)
	front := 0
	tail := 0
	queue[tail] = leftNode
	tail++

	for front < tail {
		stackOut := queue[front]
		front++

		if stackOut.Left != nil {
			queue[tail] = stackOut.Left
			tail++
		}
		if stackOut.Right != nil {
			queue[tail] = stackOut.Right
			tail++
		}

		if leftNode == stackOut {
			stackOut.Next = nil
			leftNode = queue[tail-1]
		} else {
			stackOut.Next = queue[front]
		}
	}
	return root
}

// Connect2 解法2 将二叉树看三叉树 然后进行遍历连接相近两个节点
func Connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

func traverse(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	traverse(node1.Left, node1.Right)
	traverse(node2.Left, node2.Right)
	traverse(node1.Left, node2.Right)
}
