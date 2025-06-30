package point

import "math"

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	i, j := 0, len(height)-1
	area := math.MinInt32
	for i < j {
		if height[i] < height[j] {
			temp := height[i] * (j - i)
			if temp > area {
				area = temp
			}
			i++
		} else {
			temp := height[j] * (j - i)
			if temp > area {
				area = temp
			}
			j--
		}
	}
	return area
}
