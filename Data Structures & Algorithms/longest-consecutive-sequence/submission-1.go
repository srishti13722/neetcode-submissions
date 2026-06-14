func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	if n == 0{
		return 0
	}

	curr := 1
	length := 1

	for i := 1 ; i < n ; i++{
		if nums[i] == nums[i-1]{
			continue
		}

		if nums[i] == nums[i-1]+1{
			curr++
		}else{
			curr = 1
		}

		length = max(length , curr)
	}
	

	return length
}
