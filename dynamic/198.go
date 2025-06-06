package dynamic

func rob(nums []int) int {
	var max, current int
	maxNum := make([]int, len(nums))
	dfs(nums, 0, &current, &max, maxNum)
	return max
}

func dfs(nums []int, index int, current, max *int, maxNum []int) {
	if *current > *max {
		*max = *current
	}

	for i := index; i < len(nums); i++ {
		*current += nums[i]
		if *current > maxNum[i] {
			maxNum[i] = *current
		} else {
			*current -= nums[i]
			continue
		}

		dfs(nums, i+2, current, max, maxNum)
		*current -= nums[i]
	}
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i < len(dp); i++ {
		if dp[i-1] > dp[i-2]+nums[i-1] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-2] + nums[i-1]
		}
	}
	return dp[len(dp)-1]
}

func rob3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	prev1, prev2 := 0, nums[0]
	for i := 2; i < len(nums)+1; i++ {
		if prev2 > prev1+nums[i-1] {
			prev1 = prev2
		} else {
			temp := prev1
			prev1 = prev2
			prev2 = temp + nums[i-1]
		}
	}
	return prev2
}
