package array

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := 0
	count := 0
	for _, num := range nums {
		if cur == num {
			count++
		} else if count == 0 {
			count = 1
			cur = num
		} else {
			count--
		}
	}
	return cur
}
