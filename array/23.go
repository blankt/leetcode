package array

/**
双指针法
*/

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j] == val {
			j++
		} else {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
			j++

		}
	}
	return +i
}
