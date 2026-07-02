func trap(height []int) int {
	n := len(height)
	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0] = height[0]
	for i := 1 ; i < n ; i++{
		prefix[i] = max(height[i], prefix[i-1])
	}

	suffix[n-1] = height[n-1]
	for j := n-2 ; j >= 0 ; j--{
		suffix[j] = max(height[j], suffix[j+1])
	}

	total := 0

	for i := 0 ; i< n ; i++{
		total += min(prefix[i], suffix[i]) - height[i]
	}

	return total
}
