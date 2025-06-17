package greedy

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}

	return canJumpDfs(nums, 0, dp)
}

func canJumpDfs(nums []int, index int, dp []int) bool {
	if index == len(nums)-1 {
		return true
	}

	result := dp[index]
	if result != -1 {
		return result == 1
	}

	jump := nums[index]
	if index+jump >= len(nums)-1 {
		dp[index] = 1
		return true
	}

	for j := index + 1; j <= index+jump; j++ {
		result := canJumpDfs(nums, j, dp)
		if result == true {
			dp[index] = 1
			return true
		} else {
			dp[j] = 0
		}
	}
	return false
}

func canJump2(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}

		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return false
}
