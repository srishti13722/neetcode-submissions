func lengthOfLIS(nums []int) int {
    n := len(nums)
	next := make([]int, n+1)
	

	for i := n-1 ; i >=0 ; i--{
		curr := make([]int, n+1)
		for j := i-1 ; j >= -1 ; j--{
			pick := 0
			if  j == -1 || nums[i] > nums[j]{
				pick  = 1+ next[i+1]
			}
			notpick := next[j+1]

			curr[j+1] = max(pick, notpick)
		}

		next = curr
	}


	
	return next[0]
}
