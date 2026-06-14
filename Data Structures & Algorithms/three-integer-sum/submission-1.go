func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	for i := 0 ; i < n ; i++{
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}

		j := i+1
		k := n-1

		for j < k{
			sum := nums[i] + nums[j]+ nums[k]
			if sum == 0{
				res = append(res , []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j]== nums[j-1]{
					j++
				}
				k--
				for j < k && nums[k]== nums[k+1]{
					k--
				}
			}else if sum < 0{
				j++
			}else{
				k--
			}
		}
	}

	return res
}
