func helper(left , i int, nums []int, dp []int)int{
	if i < left{
		return 0
	}

	if dp[i] != -1{
		return dp[i]
	}

	steal := nums[i] + helper(left, i-2, nums, dp)
	skip := helper(left, i-1, nums, dp)

	dp[i] = max(steal, skip)
	return dp[i]
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1{
		return nums[0]
	}

	if n == 2{
		return max(nums[0], nums[1])
	}
	
	dp1 := make([]int,n)	
	dp2 := make([]int,n)
	

	dp1[0] = nums[0]
	dp1[1] = max(nums[0], nums[1])

	dp2[1] = nums[1]
	dp2[2] = max(nums[1], nums[2])

	for i := 2 ; i < n-1 ; i++{
			steal := nums[i] + dp1[i-2]
			skip := dp1[i-1]
			dp1[i] = max(steal, skip)
	}

	for i := 3 ; i < n ; i++{
			steal := nums[i] + dp2[i-2]
			skip := dp2[i-1]
			dp2[i] = max(steal, skip)
	}
    
	return max(dp1[n-2] , dp2[n-1])
}
