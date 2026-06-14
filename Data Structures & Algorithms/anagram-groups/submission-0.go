func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for _, str := range strs{
		strChar := []byte(str)
		
		sort.Slice(strChar , func(i, j int) bool{
			return strChar[i] < strChar[j]
		})

		s := string(strChar)

		mp[s] = append(mp[s] , str)
	}

	var res [][]string

	for _ , val := range mp{
		res = append(res, val)
	}

	return res
}
