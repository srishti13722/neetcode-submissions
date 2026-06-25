func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}

	i := 0
	
	for i < len(intervals) && intervals[i][1] < newInterval[0]{
		res = append(res, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1]{
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}

	res = append(res, newInterval)

	for i < len(intervals) {
        res = append(res, intervals[i])
        i++
    }

    return res
}
