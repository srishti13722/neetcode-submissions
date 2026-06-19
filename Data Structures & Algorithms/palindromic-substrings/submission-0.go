func checkPalindrome(i, j int, s string, dp [][]int)bool{
	if i >= j{
		return true
	}

	if dp[i][j] != -1{
		if dp[i][j] == 0{
			return false
		}
		return true
	}

	if s[i] != s[j]{
		dp[i][j] = 0
		return false
	}

	isPalindrom := checkPalindrome(i+1, j-1, s, dp)
	if isPalindrom{
		dp[i][j] = 1
	}else{
		dp[i][j] = 0
	}

	if dp[i][j] == 1{
		return true
	}

	return false
}
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
		for j := range dp[i]{
			dp[i][j] = -1
		}
	}
	count := 0
    for i := 0 ; i < n ; i++{
		for j := i ; j < n ; j++{
			if checkPalindrome(i, j, s, dp){
				count++			
			}
		}
	}

	return count
}

