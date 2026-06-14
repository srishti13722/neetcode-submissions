type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder

	for _, str := range strs{
		res.WriteString(strconv.Itoa(len(str)))
		res.WriteString("#")
		res.WriteString(str)
	}

	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	n := len(encoded)
	var res []string

	i := 0

	for i < n{
		j := i

		for encoded[j] != '#'{
			j++
		}

		length , _ := strconv.Atoi(encoded[i:j])

		word := encoded[j+1 : j+1+length]

		res = append(res, word)

		i = j+1+length
	}

	return res
}
