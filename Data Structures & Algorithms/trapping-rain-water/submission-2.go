func trap(height []int) int {
	leftMax := 0
	rightMax := 0

	left := 0
	right := len(height)-1

	total := 0

	for left < right{
		if height[left] <= height[right]{
			if height[left] < leftMax{
				total += leftMax - height[left]
			}else{
				leftMax = height[left]
			}
			left++
		}else{
			if rightMax > height[right]{
				total += rightMax - height[right]
			}else{
				rightMax = height[right]
			}
			right--
		}
	}

	return total
}
