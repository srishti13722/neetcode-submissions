func lowerBound(nums []int , target int)int{
	low := 0
	high := len(nums)-1
	ans := -1
	for low <= high{
		mid := (low+high)/2
		if nums[mid] >= target{
			ans = mid
			high = mid-1
		}else{
			low = mid+1
		}
	}

	return ans
}
func lengthOfLIS(nums []int) int {
    lis := []int{}
	for i := 0 ; i < len(nums); i++{
		if len(lis) == 0{
			lis = append(lis, nums[i])
		}else{
			if lis[len(lis)-1] < nums[i]{
				lis = append(lis, nums[i])
			}else{
				lb := lowerBound(lis, nums[i])
				lis[lb] = nums[i]
			}
		}
	}

	return len(lis)
}
