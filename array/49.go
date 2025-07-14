package array

import (
	"sort"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)

		anagrams, ok := resultMap[strings.Join(chars, "")]
		if ok {
			anagrams = append(anagrams, str)
			resultMap[strings.Join(chars, "")] = anagrams
		} else {
			resultMap[strings.Join(chars, "")] = []string{str}
		}
	}

	result := make([][]string, 0, len(resultMap))
	for _, anagrams := range resultMap {
		result = append(result, anagrams)
	}
	return result
}

func groupAnagrams2(strs []string) [][]string {
	resultMap := make(map[string][]string)

	for _, str := range strs {
		indexList := make([]int, 26)
		for _, char := range str {
			index := int(char - 'a')
			indexList[index]++
		}

		key := ""
		for i := 0; i < 26; i++ {
			key += strconv.Itoa(indexList[i]) + "-"
		}

		resultMap[key] = append(resultMap[key], str)
	}

	result := make([][]string, 0, len(resultMap))
	for _, anagrams := range resultMap {
		result = append(result, anagrams)
	}

	return result
}
