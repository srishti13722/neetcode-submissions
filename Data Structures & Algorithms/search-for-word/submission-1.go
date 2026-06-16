func helper(i , j , idx int, board [][]byte, word string, visited [][]bool) bool{
	if idx == len(word)-1{
		return true
	}

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	for k := 0 ; k < 4; k++{
		r := i + dx[k]
		c := j + dy[k]

		if r >=0 && r < len(board) && c >= 0 && c < len(board[0]) && board[r][c] == word[idx+1] && !visited[r][c]{
			visited[r][c] = true
			if helper(r, c, idx+1, board, word, visited){
				return true
			}
			visited[r][c] = false
		}
	}

	return false
}
func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	for i := 0 ; i < n ; i++{
		for j := 0 ; j < m ; j++{
			if board[i][j] == word[0]{
				visited := make([][]bool, n )
				for i := range visited{
					visited[i] = make([]bool, m)
				}
				visited[i][j] = true
				if helper(i, j, 0 , board, word, visited){
					return true
				}
			}
		}
	}

	return false

}
