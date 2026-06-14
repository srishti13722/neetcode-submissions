func characterReplacement(s string, k int) int {
	n := len(s)
	left := 0
	
	ans := 0

	mp := make(map[byte]int)

	maxFreq := 0
	
	for right := 0 ; right < n ; right++{
		mp[s[right]]++
		maxFreq =  max(maxFreq, mp[s[right]])

		for right-left+1 - maxFreq > k{
			mp[s[left]]--
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
