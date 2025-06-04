package array

func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	result := make([][]string, 0)
	partitionDFS(s, 0, []string{}, &result, f)
	return result
}

func partitionDFS(s string, index int, path []string, result *[][]string, f [][]bool) {
	if index == len(s) {
		*result = append(*result, append([]string{}, path...))
		return
	}

	for j := index; j < len(s); j++ {
		subStr := s[index : j+1]
		if f[index][j] {
			path = append(path, subStr)
			partitionDFS(s, j+1, path, result, f)
			path = path[:len(path)-1]
		}
	}
}

func isValidStr(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
