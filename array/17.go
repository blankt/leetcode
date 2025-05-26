package array

import "strings"

func letterCombinations(digits string) []string {
	var numToCharMap = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	nums := make([][]string, 0, len(digits))
	for _, digit := range digits {
		if chars, ok := numToCharMap[string(digit)]; ok {
			nums = append(nums, chars)
		} else {
			return []string{} // 如果有不合法的数字，返回空切片
		}
	}

	if len(nums) == 0 {
		return []string{} // 如果没有数字，返回空切片
	}

	result := make([]string, 0)
	letterBacktrack(nums, 0, &result, []string{})

	return result
}

func letterBacktrack(nums [][]string, index int, result *[]string, track []string) {
	if index == len(nums) {
		*result = append(*result, strings.Join(track, ""))
		return
	}

	for _, char := range nums[index] {
		track = append(track, char)

		letterBacktrack(nums, index+1, result, track)

		track = track[:len(track)-1]
	}
}
