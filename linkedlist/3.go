package leetcode

func LengthOfLongestSubstring(s string) int {
	windows := make(map[string]int)

	left, right := 0, 0
	shrink := false
	lens := 0
	for right < len(s) {
		c := string(s[right])
		right++

		if times, ok := windows[c]; ok {
			times++
			if times > 1 {
				shrink = true
			}
			windows[c] = times
		} else {
			windows[c] = 1
		}

		for shrink {
			c = string(s[left])
			left++
			times, _ := windows[c]
			times--
			if times == 1 {
				shrink = false
			}
			windows[c] = times
		}

		if right-left > lens {
			lens = right - left
		}
	}

	return lens
}
