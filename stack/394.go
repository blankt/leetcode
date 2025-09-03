package stack

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	var (
		res, temp string
	)
	numStack := make([]int, 0)
	charStack := make([]string, 0)

	numTemp := ""
	for _, v := range s {
		if v >= '0' && v <= '9' {
			v = v - '0'
			numTemp = fmt.Sprintf("%v%v", numTemp, v)
		} else if v == '[' {
			if numTemp != "" {
				num, _ := strconv.Atoi(numTemp)
				numStack = append(numStack, num)
				numTemp = ""
			}
			charStack = append(charStack, string(v))
		} else if v == ']' {
			for charStack[len(charStack)-1] != "[" {
				temp = charStack[len(charStack)-1] + temp
				charStack = charStack[:len(charStack)-1]
			}
			charStack = charStack[:len(charStack)-1]

			if len(numStack) > 0 {
				multi := numStack[len(numStack)-1]
				singleTemp := temp
				for i := 0; i < multi-1; i++ {
					temp += singleTemp
				}
				numStack = numStack[:len(numStack)-1]
				charStack = append(charStack, temp)
				temp = ""
			}
		} else {
			charStack = append(charStack, string(v))
		}
	}
	for _, v := range charStack {
		res += v
	}

	return res
}
