var dx []int = []int{-1, 0, 0, 1}
var dy []int = []int{0, -1, 1, 0}
func dfs(i, j int, grid [][]byte, visited [][]bool){
	visited[i][j] = true

	for k := 0 ; k < 4 ; k++{
		r := i + dx[k]
		c := j + dy[k]

		if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == '1' && !visited[r][c]{
			dfs(r , c, grid, visited)
		}
	}
}
func numIslands(grid [][]byte) int {
    n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n)
	for i := range visited{
		visited[i] = make([]bool, m)
	}

	var count int

	for i := 0 ; i < n ; i++{
		for j := 0 ; j < m ; j++{
			if grid[i][j] == '1'&& !visited[i][j]{
				count++
				dfs(i,j,grid, visited)
			}
		}
	}

	return count
}
