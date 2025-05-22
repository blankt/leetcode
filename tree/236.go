package tree

/*
236. 二叉树的最近公共祖先
三种情况：
1. p、q在root的两侧,root是最近的公共祖先
2. q 在p的子树中则q是最近的公共祖先
3. p 在q的子树中则p是最近的公共祖先
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfsLowestAncestor(root, p, q)
}

func dfsLowestAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := dfsLowestAncestor(root.Left, p, q)
	right := dfsLowestAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}
