func helper(i int, target int, nums[]int, temp []int, ans *[][]int){
	if i == -1{
		if target == 0{		
			curr := make([]int, len(temp))
			copy(curr, temp)
			*ans = append(*ans, curr)
		}
		return
	}

	if nums[i] <= target{
		temp = append(temp, nums[i])
		helper(i, target-nums[i], nums, temp, ans)
		temp = temp[:len(temp)-1]
	}

	helper(i-1, target, nums, temp , ans)
} 

func combinationSum(nums []int, target int) [][]int {
    var ans [][]int
	var temp []int
	n := len(nums)
	helper(n-1, target, nums, temp, &ans)
	return ans
}
