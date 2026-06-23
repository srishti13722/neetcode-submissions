func helper(idx , amount int, coins []int, dp [][]int)int{
	if idx == 0{
		if amount % coins[idx] == 0{
			return amount/coins[idx]
		}

		return math.MaxInt
	}

	if dp[idx][amount] != math.MaxInt{
		return dp[idx][amount]
	}

	take := math.MaxInt
	if coins[idx] <= amount{
		coins := helper(idx, amount-coins[idx], coins, dp)
		if coins != math.MaxInt{
			take = coins +1 
		}	
	}

	nottake := helper(idx-1, amount, coins, dp)

	dp[idx][amount] = min(take , nottake)

	return dp[idx][amount]
}
func coinChange(coins []int, amount int) int {
	n := len(coins)

	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, amount+1)
		for j := range dp[i]{
			dp[i][j] = math.MaxInt
		}
	}

	for j := 0 ; j <=amount ; j++{
		if j % coins[0] == 0{
			dp[0][j] = j/coins[0]
		}
	}

	for i := 1; i < n ; i++{
		for j := 0 ; j <= amount ; j++{
			take := math.MaxInt
			if coins[i] <= j{
				coins := dp[i][j-coins[i]]
				if coins != math.MaxInt{
					take = coins +1 
				}	
			}
			nottake := dp[i-1][j]

			dp[i][j] = min(take , nottake)
		}
	}


    
	nums := dp[n-1][amount]

	if nums == math.MaxInt{
		return -1
	}

	return nums
}
