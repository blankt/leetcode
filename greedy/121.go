package greedy

import "math"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	result := 0
	min := math.MaxInt32

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i] > min && prices[i]-min > result {
			result = prices[i] - min
		}
	}
	return result
}
