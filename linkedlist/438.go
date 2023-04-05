package leetcode

func findAnagrams(s string, p string) []int {
	windows := make(map[string]int)
	needs := make(map[string]int)
	for _, v := range p {
		times := needs[string(v)]
		times++
		needs[string(v)] = times
	}

	left, right := 0, 0
	needSize := 0
	result := make([]int, 0)

	for right < len(s) {
		c := string(s[right])
		right++

		if times, ok := needs[c]; ok {
			hadTimes := windows[c]
			hadTimes++
			if hadTimes == times {
				needSize++
			}
			windows[c] = hadTimes
		}

		for needSize >= len(needs) {
			if right-left == len(p) {
				result = append(result, left)
			}
			c = string(s[left])
			left++

			if times, ok := needs[c]; ok {
				hadTimes := windows[c]
				hadTimes--
				if hadTimes < times {
					needSize--
				}
				windows[c] = hadTimes
			}
		}
	}
	return result
}
