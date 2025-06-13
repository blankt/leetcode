package dynamic

/*
相当于求在nums数组中是否有子集相加等于数组和的一半
子问题：求在nums[0:i]中是否有子集相加等于j
状态转移方程：f（j） = f(j) || f(j - nums[i])
*/
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		x := make([]int, target+1)
		for j := 0; j < len(x); j++ {
			x[j] = -1
		}
		dp[i] = x
	}

	var canPartitionDfs func(i, j int) bool
	canPartitionDfs = func(i, j int) bool {
		if i < 0 {
			return j == 0
		}

		p := dp[i][j]
		if p != -1 {
			return p == 1
		}

		res := j >= nums[i] && canPartitionDfs(i-1, j-nums[i]) || canPartitionDfs(i-1, j)
		if res {
			dp[i][j] = 1
			return true
		} else {
			dp[i][j] = 0
			return false
		}
	}

	return canPartitionDfs(len(nums)-1, target)
}
