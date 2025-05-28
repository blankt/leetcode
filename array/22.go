package array

import "strings"

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	uniqueMap := make(map[string]struct{})
	charNumMap := make(map[string]int)
	parenthesesTrackback([]string{"(", ")"}, []string{}, &result, uniqueMap, charNumMap, n*2, 0)
	return result
}

func parenthesesTrackback(parentheses []string, trace []string, result *[]string, uniqueMap map[string]struct{}, charNumMap map[string]int, n, start int) {
	if len(trace) == n {
		resultStr := strings.Join(trace, "")
		_, exists := uniqueMap[resultStr]

		if isValid(resultStr) && !exists {
			uniqueMap[resultStr] = struct{}{}
			*result = append(*result, strings.Join(trace, ""))
		}
		return
	}

	for i := start; i < n; i++ {
		for _, char := range parentheses {
			if len(trace) > 0 && string(trace[0]) == ")" {
				continue
			}
			if charNumMap[char] >= n/2 {
				continue
			}

			trace = append(trace, char)
			charNumMap[char] += 1
			parenthesesTrackback(parentheses, trace, result, uniqueMap, charNumMap, n, i+1)
			trace = trace[:len(trace)-1]
			charNumMap[char] -= 1
		}
	}
}

func isValid(parentheses string) bool {
	if len(parentheses)%2 != 0 || string(parentheses[0]) == ")" {
		return false
	}
	stack := make([]rune, 0)
	for _, char := range parentheses {
		if string(char) == "(" {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if string(top) != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
