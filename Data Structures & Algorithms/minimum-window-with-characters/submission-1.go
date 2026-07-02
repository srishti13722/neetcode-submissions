func minWindow(s string, t string) string {
    n := len(s)
	m := len(t)

	if n < m {
		return ""
	}

	freq := make(map[byte]int)
	for i := 0 ; i< m ; i++{
		freq[t[i]]++
	}

	left := 0
	countReq := len(t)
	minLen := math.MaxInt
	ansi := 0

	for right := 0 ; right <n ; right++{
		if freq[s[right]] > 0{
			countReq--
			for countReq == 0{
				if right-left+1 < minLen{
					minLen = right-left+1
					ansi = left
				}
				freq[s[left]]++
				if freq[s[left]] > 0{
					countReq++
				}
				left++
			}
		}
		
		freq[s[right]]--
		
	}

	if minLen == math.MaxInt{
		return ""
	}

	return s[ansi:ansi+minLen]
}
