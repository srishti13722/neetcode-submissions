func mergeInterval(intervals [][]int) [][]int{
	var res [][]int
	res = append(res, intervals[0])
	
	

	for i := 1 ; i < len(intervals) ; i++{
		prevEnd := res[len(res)-1][1]

		currStart := intervals[i][0]

		if prevEnd >= currStart{
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		}else{
			res = append(res, intervals[i])
		}
	}

	return res
}
func insert(intervals [][]int, newInterval []int) [][]int {
    intervals = append(intervals, newInterval)
	
	 sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

	return mergeInterval(intervals)
}
