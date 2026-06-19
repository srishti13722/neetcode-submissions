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

	if n == 1{
		return nums[0]
	}
    
	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2 ; i < n ; i++{
		steal := nums[i] + prev2
		skip := prev1
		curr := max(steal, skip)
		prev2 = prev1
		prev1 = curr
	}
	
	return prev1
}
