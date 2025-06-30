package point

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
}
