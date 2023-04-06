package array

import "math"

func MinSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := math.MaxInt
	i := 0
	j := 0
	sum := 0
	for j < len(nums) {
		sum += nums[j]

		if sum >= target {
			for sum >= target {
				if j-i+1 < n {
					n = j - i + 1
				}
				sum -= nums[i]
				i++
			}
		}
		j++
	}
	if n == math.MaxInt {
		return 0
	} else {
		return n
	}
}
