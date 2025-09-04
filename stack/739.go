package stack

/*
73,74,75,71,69,72,76,73
1,1,4,2,1,1,0,0
*/

func dailyTemperatures2(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, v := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
		}
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}

func dailyTemperatures1(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := len(temperatures) - 2; i >= 0; i-- {
		step := 1
		exist := false
		for j := i + 1; j <= len(temperatures)-1; {
			if temperatures[j] > temperatures[i] {
				exist = true
				break
			} else if result[j] == 0 {
				step++
				j++
				break
			} else {
				step += result[j]
				j += result[j]
			}
		}
		if exist {
			result[i] = step
		}
	}
	return result
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i, v := range temperatures {
		exist := false
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > v {
				result[i] = j - i
				exist = true
				break
			}
		}
		if !exist {
			result[i] = 0
		}
	}
	return result
}
