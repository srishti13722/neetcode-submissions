func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1

	for i := 1 ; i < len(nums) ; i++{
		res[i] = nums[i-1] * res[i-1]
	}

	suffix := 1
	res[len(res)-1] = res[len(nums)-1] * suffix

	for j := len(nums)-2; j >= 0 ; j--{
		curr := nums[j+1] * suffix
		res[j] = res[j] * curr
		suffix = curr
	}

	return res
}
