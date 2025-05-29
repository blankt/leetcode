package array

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existDFS(board, word, i, j, 0) {
				return true
			}

		}
	}

	return false
}

func existDFS(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}

	if len(word)-1 == k {
		return true
	}

	board[i][j] = ' '

	res := existDFS(board, word, i-1, j, k+1) || existDFS(board, word, i+1, j, k+1) || existDFS(board, word, i, j+1, k+1) || existDFS(board, word, i, j-1, k+1)

	board[i][j] = word[k]

	return res
}
