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
	dp1 := make([]int,n)
	for i := range dp1{
		dp1[i] = -1
	}

	dp2 := make([]int,n)
	for i := range dp2{
		dp2[i] = -1
	}
    return max(helper(0,n-2, nums, dp1) , helper(1, n-1, nums, dp2))
}
