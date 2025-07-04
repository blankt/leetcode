package leetcode

func sortList(head *ListNode) *ListNode {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	nums = mergeSort(nums)

	dummy := &ListNode{}
	dummyHead := dummy
	for _, num := range nums {
		dummyHead.Next = &ListNode{Val: num}
		dummyHead = dummyHead.Next
	}
	return dummy.Next
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, append(left[i:], right[j:]...)...)
	return result
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}
