func trap(height []int) int {
	n := len(height)
	prefix := make([]int, n)

	prefix[0] = height[0]
	for i := 1 ; i < n ; i++{
		prefix[i] = max(height[i], prefix[i-1])
	}

	suffix := height[n-1]
	total := 0

	for i := n-2 ; i >= 0 ; i--{
		suffix = max(suffix , height[i])
		total += min(prefix[i], suffix) - height[i]		
	}

	return total
}
