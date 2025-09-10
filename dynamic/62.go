package dynamic

func uniquePaths(m int, n int) int {
	// start 	0,0
	// end 		m-1,n-1
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	var result int
	dfsUniquePaths(dp, 0, 0, &result)
	return result
}

var directions = [][]int{
	{0, 1}, // right
	{1, 0}, // down
}

func dfsUniquePaths(visited [][]bool, x, y int, result *int) {
	if x < 0 || y < 0 || x >= len(visited) || y >= len(visited[0]) || visited[x][y] {
		return
	}

	if x == len(visited)-1 && y == len(visited[0])-1 {
		*result++
		return
	}

	visited[x][y] = true
	for _, dir := range directions {
		dfsUniquePaths(visited, x+dir[0], y+dir[1], result)
	}
	visited[x][y] = false
}

// 动态规划
// 状态转移方程 dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePaths1(m int, n int) int {
	// start 	0,0
	// end 		m-1,n-1
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
