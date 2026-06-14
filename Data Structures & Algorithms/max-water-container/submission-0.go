func maxArea(heights []int) int {
	n := len(heights)
	low := 0
	high := n-1

	ans := 0

	for low < high{
		curr := min(heights[high], heights[low]) * (high-low)
		ans = max(ans, curr)

		if heights[high] < heights[low]{
			high--
		}else{
			low++
		}
	}

	return ans
}
