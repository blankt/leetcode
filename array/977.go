package array

// 双指针法
func sortedSquares(nums []int) []int {
	x := 0
	for x < len(nums) {
		if nums[x] > 0 {
			break
		}
		x++
	}

	result := make([]int, 0, len(nums))
	i := x - 1
	for i >= 0 && x < len(nums) {
		var a int
		if nums[i] < 0 {
			a = -nums[i]
		} else {
			a = nums[i]
		}

		if a < nums[x] {
			result = append(result, a*a)
			i--
		} else {
			result = append(result, nums[x]*nums[x])
			x++
		}
	}

	for i >= 0 {
		result = append(result, nums[i]*nums[i])
		i--
	}

	for x < len(nums) {
		result = append(result, nums[x]*nums[x])
		x++
	}

	return result
}
