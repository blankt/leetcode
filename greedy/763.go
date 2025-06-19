package greedy

func partitionLabels(s string) []int {
	if len(s) == 0 {
		return nil
	}
	result := make([]int, 0)
	charNumMap := make(map[uint8]int)
	for i := range s {
		c := s[i]
		num, ok := charNumMap[c]
		if !ok {
			charNumMap[c] = 1
		} else {
			charNumMap[c] = num + 1
		}
	}

	start, end := 0, 0
	existedNumMap := make(map[uint8]struct{})
	for end < len(s) {
		c := s[end]
		existedNumMap[c] = struct{}{}
		charNumMap[c]--

		existChar := false
		for k, _ := range existedNumMap {
			if num, ok := charNumMap[k]; ok {
				if num > 0 {
					existChar = true
					// 如果当前字符还有剩余，则需要继续往后找
					break
				} else {
					// 如果当前字符已经没有了，则从已存在的字符中删除
					delete(existedNumMap, k)
				}
			}
		}

		if !existChar {
			result = append(result, end-start+1)
			existedNumMap = make(map[uint8]struct{})
			start, end = end+1, end+1
		} else {
			end++
		}
	}
	return result
}

func partitionLabels1(s string) []int {
	charLastIndex := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		charLastIndex[s[i]] = i
	}

	result := make([]int, 0)
	start, end := 0, 0
	for i := range s {
		indexEnd := charLastIndex[s[i]]
		if indexEnd > end {
			end = indexEnd
		}

		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
