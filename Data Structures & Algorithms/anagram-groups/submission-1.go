func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)

	for _, str := range strs{
		var freq [26]int

		for i := 0 ; i< len(str); i++{
			freq[str[i]-'a']++
		}

		mp[freq] = append(mp[freq], str)
	}

	var res [][]string

	for _, val := range mp{
		res = append(res, val)
	}

	return res
}
