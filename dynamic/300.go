package dynamic

/*
 */
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && max < dp[j] {
				max = dp[j]
			}
		}
		dp[i] = max + 1
	}

	max := -1
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
