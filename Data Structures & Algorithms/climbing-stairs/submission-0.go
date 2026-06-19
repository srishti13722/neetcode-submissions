func helper(n int, dp []int)int{
	if n == 0 || n == 1 || n == 2{
		return n
	}

	if dp[n] != -1{
		return dp[n]
	}

	dp[n] = helper(n-1,dp) + helper(n-2,dp)

	return dp[n]
}
func climbStairs(n int) int {
	if n <= 2 {
        return n
    }
	
	dp := make([]int, n+1)
	for i := range dp{
		dp[i] = -1
	}
    
	return helper(n, dp)
}
