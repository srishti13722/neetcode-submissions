func findMin(nums []int) int {
	n := len(nums)
	low := 0
	high := n-1

	mini := math.MaxInt

	for low <= high{
		mid := low + (high - low)/2

		if nums[low] <= nums[mid]{
			mini = min(mini, nums[low])
			low = mid+1
		}else{
			mini = min(mini, nums[mid])
			high = mid-1		
		}
	}

	return mini
}
