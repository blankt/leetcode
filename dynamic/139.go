package dynamic

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	for _, v := range wordDict {
		if strings.HasPrefix(s, v) {
			if len(s) == len(v) {
				return true
			}

			result := wordBreak(strings.TrimPrefix(s, v), wordDict)
			if result {
				return true
			}
		}
		continue
	}
	return false
}

func wordBreak1(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
