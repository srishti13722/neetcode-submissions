func dfs(node int, graph map[int][]int, visited map[int]bool){
    visited[node] = true

    for _, ngbr := range graph[node]{
        if !visited[ngbr]{
            dfs(ngbr, graph, visited)
        }
    }
}
func countComponents(n int, edges [][]int) int {
   graph := make(map[int][]int)

   for _, edge := range edges{
        u := edge[0]
        v := edge[1]

        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
   } 

   count := 0

   visited := make(map[int]bool)

   for i := 0 ; i < n; i++{
        if !visited[i]{
            count++
            dfs(i, graph, visited)
        }
   }

   return count
}
