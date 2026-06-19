func helper(i int, nums []int, dp []int)int{
	if i < 0{
		return 0
	}

	if dp[i] != -1{
		return dp[i]
	}

	steal := nums[i] + helper(i-2, nums, dp)
	skip := helper(i-1, nums, dp)

	dp[i] = max(steal, skip)
	return dp[i]
}

func rob(nums []int) int {
	n := len(nums)
    dp := make([]int, n)
	for i := range dp{
		dp[i] = -1
	}
	
	return helper(n-1, nums, dp)
}
