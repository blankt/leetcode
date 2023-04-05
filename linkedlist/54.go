package leetcode

func SpiralOrder(matrix [][]int) []int {
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	up := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	size := 0
	num := len(matrix) * len(matrix[0])
	for size < num {
		if left <= right {
			i := left
			for i <= right {
				result = append(result, matrix[up][i])
				i++
				size++
			}
			up++
		}

		if size >= num {
			break
		}

		if up <= down {
			i := up
			for i <= down {
				result = append(result, matrix[i][right])
				i++
				size++
			}
			right--
		}

		if size >= num {
			break
		}

		if right >= left {
			i := right
			for i >= left {
				result = append(result, matrix[down][i])
				i--
				size++
			}
			down--
		}
		if size >= num {
			break
		}
		if down >= up {
			i := down
			for i >= up {
				result = append(result, matrix[i][left])
				i--
				size++
			}
			left++
		}
	}
	return result
}
