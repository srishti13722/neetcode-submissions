func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)

	left := 0
	n := len(s)

	ans := 0

	for right := 0 ; right < n; right++{
		if idx , ok := mp[s[right]] ; ok && idx >= left{
			left = idx + 1
		}

		ans = max(ans, right-left+1)
		mp[s[right]] = right

	}

	return ans
}
