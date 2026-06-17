var dx []int = []int{-1, 0, 0, 1}
var dy []int = []int{0, -1, 1, 0}

func dfs(i, j int, heights [][]int, pacific [][]bool){
	pacific[i][j] = true

	for k := 0 ; k < 4; k++{
		r := i + dx[k]
		c := j + dy[k]

		if r >= 0 && r < len(heights) && c >= 0 && c < len(heights[0]){
			if heights[i][j] <= heights[r][c] && pacific[r][c] == false{
				dfs(r, c, heights, pacific)
			}
		}
	}
}
func pacificAtlantic(heights [][]int) [][]int {
    n := len(heights)
	m := len(heights[0])

	pacific := make([][]bool , n)
	for i := range pacific{
		pacific[i] = make([]bool, m)
	}

	atlantic := make([][]bool , n)
	for i := range atlantic{
		atlantic[i] = make([]bool, m)
	}

	for j := 0; j < m; j++ {
    if !pacific[0][j] {
        dfs(0, j, heights, pacific)
    }
    if !atlantic[n-1][j] {
        dfs(n-1, j, heights, atlantic)
    }
}

for i := 0; i < n; i++ {
    if !pacific[i][0] {
        dfs(i, 0, heights, pacific)
    }
    if !atlantic[i][m-1] {
        dfs(i, m-1, heights, atlantic)
    }
}

	var res [][]int

	for i := 0 ; i < n ; i++{
		for j := 0 ; j < m ; j++{
			if pacific[i][j] && atlantic[i][j]{
				res = append(res, []int{i,j})
			}
		}
	}

	return res
}
