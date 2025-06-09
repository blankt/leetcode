package dynamic

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if minn > dp[i-j*j] {
				minn = dp[i-j*j]
			}
		}
		dp[i] = minn + 1
	}
	return dp[n]
}
