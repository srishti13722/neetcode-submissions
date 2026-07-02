func checkInclusion(s1 string, s2 string) bool {
	freqS1 := [26]int{}
	for i := 0 ; i < len(s1) ; i++{
		freqS1[s1[i]-'a']++
	}

	left := 0
	temp := [26]int{}
	for right := 0 ; right < len(s2); right++{
		temp[s2[right]-'a']++
		
		if right-left+1 > len(s1){
			temp[s2[left]-'a']--
			left++
		}

		if freqS1 == temp{
			return true
		}		
	}

	return false
}
