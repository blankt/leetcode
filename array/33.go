package array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	index := 0
	for index < len(nums)-2 {
		if nums[index] > nums[index+1] {
			break
		}
		index++
	}

	newNums := append(nums[index+1:], nums[:index+1]...)
	result := binarySearch(newNums, target)
	if result == -1 {
		return result
	} else {
		return result + index + 1
	}
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
