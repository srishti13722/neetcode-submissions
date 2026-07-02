func maxSlidingWindow(nums []int, k int) []int {
    var res []int
	n := len(nums)

	for i := 0 ; i <= n-k; i++{
		maxi := nums[i]
		for j := i ; j < i + k ; j++{
			if nums[j] >maxi{
				maxi = nums[j]
			} 
		}

		res = append(res, maxi)
	}

	return res
}
