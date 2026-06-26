func lengthOfLIS(nums []int) int {
    n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}

	for i := n-1 ; i >=0 ; i--{
		for j := i-1 ; j >= -1 ; j--{
			pick := 0
			if  j == -1 || nums[i] > nums[j]{
				pick  = 1+ dp[i+1][i+1]
			}
			notpick := dp[i+1][j+1]

			dp[i][j+1] = max(pick, notpick)
		}
	}


	
	return dp[0][0]
}
