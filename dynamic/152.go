package dynamic

import "math"

func maxProduct(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := nums[i]
			tempMax := nums[i]
			for j := i - 1; j >= 0; j-- {
				temp *= nums[j]
				if temp > tempMax {
					tempMax = temp
				}
			}
			dp[i] = tempMax
		} else {
			if nums[i]*dp[i-1] > nums[i] {
				dp[i] = nums[i] * dp[i-1]
			} else {
				dp[i] = nums[i]
			}
		}
	}

	maxNum := math.MinInt16
	for i := 0; i < len(nums); i++ {
		if dp[i] > maxNum {
			maxNum = dp[i]
		}
	}
	return maxNum
}
