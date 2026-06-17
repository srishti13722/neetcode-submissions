type DisjointSet struct{
	rank []int
	parent []int
}

func NewDisjointSet(n int)*DisjointSet{
	rank := make([]int, n+1)
	parent := make([]int, n+1)
	for i := range parent{
		parent[i] = i
	}

	return &DisjointSet{
		rank : rank,
		parent : parent,
	}
}

func(d *DisjointSet)UnionByRank(u, v int){
	pu := d.FindParent(u)
	pv := d.FindParent(v)

	if pu == pv{
		return
	}

	rpu := d.rank[pu]
	rpv := d.rank[pv]

	if rpu > rpv{
		d.parent[pv] = pu
	}else if rpv > rpu{
		d.parent[pu] = pv
	}else{
		d.rank[pu]++
		d.parent[pv] = pu
	}
}

func(d *DisjointSet)FindParent(u int)int{
	if d.parent[u] == u{
		return u
	}

	d.parent[u] = d.FindParent(d.parent[u])
	return d.parent[u]
}

func countComponents(n int, edges [][]int) int {
    ds := NewDisjointSet(n)

	for _, edge := range edges{
		u := edge[0]
		v := edge[1]

		ds.UnionByRank(u,v)
	}

	count := 0

	for i := 0 ; i < n; i++{
		if ds.FindParent(i) == i{
			count++
		}
	}

	return count
}
