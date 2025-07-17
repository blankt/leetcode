package array

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		if nums[left] == target {
			index = left
			break
		}
		if nums[right] == target {
			index = right
			break
		}

		mid := (left + right) / 2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if index == -1 {
		return []int{-1, -1}
	}

	left, right = 0, index-1
	otherTarget := float64(target) - 0.5
	for left <= right {
		mid := (left + right) / 2
		if float64(nums[mid]) < otherTarget {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	start := left

	otherTarget = float64(target) + 0.5
	left, right = index+1, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if float64(nums[mid]) > otherTarget {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	end := right

	return []int{start, end}
}
