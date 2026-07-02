func sortString(str string)string{
	runes := []rune(str)
	sort.Slice(runes, func(i, j int)bool{
		return runes[i] < runes[j]
	})

	return string(runes)
}

func checkSame(s1 string, s2 string)bool{
	if sortString(s1) == sortString(s2){
		return true
	}

	return false
}

func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)

	for i := 0 ; i <= m-n ; i++{
		if checkSame(s1, s2[i:i+n]){
			return true
		}
	}

	return false
}
