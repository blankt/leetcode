package dynamic

import "math"

/*
1. 找出子问题 请最长连续子序列的最大乘积 等于求子序列的i-1连续序列的最大值
2. 找出状态转移方程 f(x) = max(f(x),f(x-1) * nums[x])
3. 由于存在负数的情况 会导致原来的最小值变成最大值 所以需要维护请求前面序列的最小值 当为负数的时候计算最大值
*/
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

func maxProduct2(nums []int) int {
	maxNum := math.MinInt16
	imax := 1
	imin := 1

	for _, v := range nums {
		if v < 0 {
			temp := imax
			imax = imin
			imin = temp
		}

		imax = max(imax*v, v)
		imin = min(imin*v, v)
		maxNum = max(maxNum, imax)
	}

	return maxNum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
