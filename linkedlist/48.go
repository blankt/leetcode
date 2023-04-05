package leetcode

func Rotate(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	for i := 0; i < len(matrix); i++ {
		left := 0
		right := len(matrix) - 1
		for left < right {
			temp := matrix[i][left]
			matrix[i][left] = matrix[i][right]
			matrix[i][right] = temp
			left++
			right--
		}
	}

	return matrix
}
