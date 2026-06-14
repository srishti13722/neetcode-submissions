func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0] = 1
	suffix[len(nums)-1] = 1

	for i := 1 ; i < len(nums) ; i++{
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	for j := len(nums)-2; j >= 0 ; j--{
		suffix[j] = nums[j+1] * suffix[j+1]
	}

	res := make([]int, len(nums))

	for i := 0 ; i< len(res) ; i++{
		res[i] = prefix[i] * suffix[i]
	}

	return res
}
