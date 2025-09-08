package array

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] <= nums[j] {
			for i >= 0 {
				if i-1 < 0 {
					return nums[i]
				}
				if nums[i] < nums[i-1] {
					return nums[i]
				}
				i--
			}
		}
		mid := (i + j) / 2
		i = mid + 1
	}

	return -1
}

// 二分查找精髓在于不断的缩小左右边界，直至找到最大和最小的数子
func findMin1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		}
	}

	return nums[left]
}
