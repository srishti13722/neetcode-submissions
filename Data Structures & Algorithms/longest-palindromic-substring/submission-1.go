func checkPalindrome(i, j int, s string)bool{
	if i >= j{
		return true
	}

	if s[i] != s[j]{
		return false
	}

	return checkPalindrome(i+1, j-1, s)
}
func longestPalindrome(s string) string {
	n := len(s)
	maxi := 0
	longest := ""
    for i := 0 ; i < n ; i++{
		for j := i ; j < n ; j++{
			if checkPalindrome(i, j, s){
				if j - i +1 > maxi{
					maxi = j-i+1 
					longest = s[i:j+1]
				}				
			}
		}
	}

	return longest
}
