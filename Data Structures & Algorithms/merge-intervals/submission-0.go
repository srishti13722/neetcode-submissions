func merge(intervals [][]int) [][]int {
    sort.Slice(intervals , func(i, j int)bool{
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	res = append(res, intervals[0])

	for i := 1 ; i< len(intervals) ; i++{
		last := res[len(res)-1][1]
		curr := intervals[i][0]

		if last >= curr{
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}else{
			res = append(res, intervals[i])
		}
	}

	return res
}
