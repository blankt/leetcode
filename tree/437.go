package tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PathSum 437. 路径总和 III
// 层序+后序遍历 计算每一颗子树是否存在这种值
func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	sum := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		front := queue.Front().Value.(*TreeNode)

		// 计算单个树
		currentSum := 0
		singleTreeSum(front, targetSum, &currentSum, &sum)

		if front.Left != nil {
			queue.PushBack(front.Left)
		}
		if front.Right != nil {
			queue.PushBack(front.Right)
		}
		queue.Remove(queue.Front())
	}
	return sum
}

func singleTreeSum(root *TreeNode, targetSum int, currentSum *int, sum *int) {
	if root == nil {
		return
	}
	*currentSum += root.Val

	singleTreeSum(root.Left, targetSum, currentSum, sum)
	singleTreeSum(root.Right, targetSum, currentSum, sum)

	if *currentSum == targetSum {
		*sum++
	}
	*currentSum -= root.Val
}

// PathSum2 主函数：路径总和等于 targetSum 的路径数量
func PathSum2(root *TreeNode, targetSum int) int {
	prefixSum := map[int]int{0: 1}
	return dfs(root, 0, targetSum, prefixSum)
}

// DFS + 前缀和 + 回溯
func dfs(node *TreeNode, currSum, targetSum int, prefixSum map[int]int) int {
	if node == nil {
		return 0
	}

	currSum += node.Val
	count := prefixSum[currSum-targetSum]

	// 添加当前节点的前缀和
	prefixSum[currSum]++

	count += dfs(node.Left, currSum, targetSum, prefixSum)
	count += dfs(node.Right, currSum, targetSum, prefixSum)

	// 回溯，移除当前节点的前缀和
	prefixSum[currSum]--
	return count
}
