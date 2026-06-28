func lengthOfLIS(nums []int) int {
    lis := make([]int, len(nums))
	for i := range lis{
		lis[i] = 1
	}

	maxi := 1

	for i := 0 ; i < len(nums) ; i++{
		for j := 0 ; j < i ; j++{
			if nums[i] > nums[j]{
				lis[i] = max(lis[i], 1 + lis[j])
				maxi = max(maxi, lis[i])
			}
		}
	}

	return maxi
}
