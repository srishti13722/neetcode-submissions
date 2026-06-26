func helper(idx , lastPicked int, nums []int, dp [][]int)int{
	if idx == len(nums){
		return 0
	}

	if dp[idx][lastPicked+1] != -1{
		return dp[idx][lastPicked+1]
	}

	pick := 0
	if lastPicked == -1 || nums[idx] > nums[lastPicked]{
		pick  = 1+ helper(idx+1, idx, nums, dp)
	}
	notpick := helper(idx+1, lastPicked, nums, dp)

	dp[idx][lastPicked+1] = max(pick, notpick)
	
	return dp[idx][lastPicked+1]
}

func lengthOfLIS(nums []int) int {
    n := len(nums)
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n+1)
		for j := range dp[i]{
			dp[i][j] = -1
		}
	}
	
	return helper(0, -1, nums, dp)
}
