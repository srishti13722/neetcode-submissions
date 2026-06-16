func helper(i , j , idx int, board [][]byte, word string) bool{
	if idx == len(word)-1{
		return true
	}

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	original := board[i][j]
	board[i][j] = '#'

	for k := 0 ; k < 4; k++{
		r := i + dx[k]
		c := j + dy[k]

		if r >=0 && r < len(board) && c >= 0 && c < len(board[0]) && board[r][c] == word[idx+1]{
			if helper(r, c, idx+1, board, word){
				board[i][j] = original
				return true
			}
		}
	}

	board[i][j] = original
	return false
}
func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	for i := 0 ; i < n ; i++{
		for j := 0 ; j < m ; j++{
			if board[i][j] == word[0]{
				if helper(i, j, 0 , board, word){
					return true
				}
			}
		}
	}

	return false

}
func findWords(board [][]byte, words []string) []string {
    var res []string

	for _, word := range words{
		if exist(board, word){
			res = append(res, word)
		}
	}

	return res
}
