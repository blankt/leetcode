package leetcode

import (
	"math"
)

func CheckInclusion(s1 string, s2 string) bool {
	needs := make(map[string]int)
	windows := make(map[string]int)
	for _, v := range s1 {
		num := needs[string(v)]
		num++
		needs[string(v)] = num
	}

	right, left, needSize := 0, 0, 0
	lens := math.MaxInt
	start := 0
	for right < len(s2) {
		c := string(s2[right])
		right++

		num := windows[c]
		num++
		if num2, ok := needs[c]; ok {
			if num == num2 {
				needSize++
			}
		}
		windows[c] = num

		for needSize >= len(needs) {
			if right-left < lens {
				start = left
			}
			if right-start == len(s1) {
				return true
			}
			leftC := string(s2[left])
			left++
			num2 := windows[leftC]
			num2--
			if nums, ok := needs[leftC]; ok {
				if num2 < nums {
					needSize--
				}
			}
			windows[leftC] = num2
		}
	}
	return false
}
