func topoSort(numCourses int, graph map[int][]int, indegree []int) []int{
	// visited := make([]bool, numCourses)
	
	topo := []int{}
	queue := []int{}
	
	for i := range indegree{
		if indegree[i] == 0{
			queue = append(queue, i)
		}
	}

	for len(queue) > 0{
		node := queue[0]
		queue = queue[1:]

		topo = append(topo, node)

		for _, ngbr := range graph[node]{
			indegree[ngbr]--
			if indegree[ngbr]==0{
				queue = append(queue, ngbr)
			}
		}
	}
	
	return topo
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	indegree := make([]int, numCourses)
	for _, edge := range prerequisites{
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
		indegree[v]++
	}
    topo := topoSort(numCourses, graph, indegree)
	if len(topo) == numCourses{
		return true
	}
	return false
}
