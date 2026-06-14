func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))

	prefix[0] = 1

	res := make([]int, len(nums))

	for i := 1 ; i < len(nums) ; i++{
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	suffix := 1
	res[len(res)-1] = prefix[len(nums)-1] * suffix

	for j := len(nums)-2; j >= 0 ; j--{
		curr := nums[j+1] * suffix
		res[j] = prefix[j] * curr
		suffix = curr
	}

	return res
}
