func robRange(nums []int)int{
	if len(nums) == 1{
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2 ; i < len(nums); i++{
		curr := max(nums[i]+prev2, prev1)
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
func rob(nums []int) int {
    n := len(nums)

	if n ==1{
		return nums[0]
	}

	ans1 := robRange(nums[:n-1])
	ans2 := robRange(nums[1:])

	return max(ans1, ans2)
}
