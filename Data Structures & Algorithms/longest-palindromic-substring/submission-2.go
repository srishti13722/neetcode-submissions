func checkPalindrome(i, j int, s string, dp [][]bool)bool{
	if i >= j{
		return true
	}

	if dp[i][j] != false{
		return dp[i][j]
	}

	if s[i] != s[j]{
		return false
	}

	dp[i][j] = checkPalindrome(i+1, j-1, s, dp)

	return dp[i][j]
}
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp{
		dp[i] = make([]bool, n)
	}
	maxi := 0
	longest := ""
    for i := 0 ; i < n ; i++{
		for j := i ; j < n ; j++{
			if checkPalindrome(i, j, s, dp){
				if j - i +1 > maxi{
					maxi = j-i+1 
					longest = s[i:j+1]
				}				
			}
		}
	}

	return longest
}
