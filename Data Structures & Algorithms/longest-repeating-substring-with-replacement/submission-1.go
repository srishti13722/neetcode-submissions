func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	maxFreq := 0

	left := 0

	maxLen := 0

	for right := 0 ; right< len(s); right++{
		freq[s[right]]++
		maxFreq = max(maxFreq, freq[s[right]])

		if right-left+1 - maxFreq > k{
			freq[s[left]]--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
