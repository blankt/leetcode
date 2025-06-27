package array

/*
[1,2,3,4]
[1,1,2,6]
[24,12,4,1]
[24,12,8,6]
*/
func productExceptSelf(nums []int) []int {
	l, r := make([]int, len(nums)), make([]int, len(nums))
	l[0], r[len(nums)-1] = 1, 1

	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = l[i] * r[i]
	}
	return result
}
