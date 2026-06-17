func dfs(node int, graph map[int][]int, visited []bool){
	visited[node] = true

	for _, ngbr := range graph[node]{
		if !visited[ngbr]{
			dfs(ngbr, graph, visited)
		}
	}

}
func validTree(n int, edges [][]int) bool {
    if len(edges) != n-1{
		return false
	}

	graph := make(map[int][]int)
	for _, edge := range edges{
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)

	dfs(0, graph, visited) 

	for i := range visited{
		if visited[i] == false{
			return false
		}
	}

	return true
}
