package array

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	x := 0
	y := len(matrix[0]) - 1
	for x < len(matrix) {
		if matrix[x][0] > target {
			return false
		}
		if matrix[x][y] < target {
			x++
			continue
		}

		start, end := 0, y
		for start <= end {
			if matrix[x][start] == target || matrix[x][end] == target {
				return true
			}
			mid := (start + end) / 2
			if matrix[x][mid] == target {
				return true
			} else if matrix[x][mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		return false
	}

	return false
}
