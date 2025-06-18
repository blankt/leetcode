package greedy

import (
	"math"
)

func jump(nums []int) int {
	dp := make([]int, len(nums))

	for i := range nums {
		minStep := math.MaxInt32
		for j := 0; j < i; j++ {
			if j+nums[j] >= i && dp[j]+1 < minStep {
				minStep = dp[j] + 1
			}
		}
		if minStep == math.MaxInt32 {
			dp[i] = 0
		} else {
			dp[i] = minStep
		}
	}
	return dp[len(nums)-1]
}

func jump2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	step := 0
	finished := false
	left, right := 0, 1

	for !finished {
		step++
		maxIndex := math.MinInt32
		for left < right {
			if left+nums[left] > maxIndex {
				maxIndex = left + nums[left]
			}

			if maxIndex >= len(nums)-1 {
				finished = true
				break
			}
			left++
		}

		left = right
		right = maxIndex + 1
	}

	return step
}
