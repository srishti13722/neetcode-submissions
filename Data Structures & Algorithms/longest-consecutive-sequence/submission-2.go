func longestConsecutive(nums []int) int {
	n := len(nums)

	if n == 0{
		return 0
	}

	mp := make(map[int]bool)

	for _ , num := range nums{
		mp[num] = true
	}

	longest := 0

	for _, num := range nums{
		if _, ok := mp[num-1] ; ok{
			continue
		}

		length := 1
		curr := num

		for mp[curr+1]{
			length++
			curr++
		}

		longest = max(longest , length)

	}

	return longest
}
