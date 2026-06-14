func maxProfit(prices []int) int {
	ans := 0
	bestbuy := prices[0]

	for i := 1 ; i < len(prices) ; i++{
		curr := prices[i] -  bestbuy
		ans = max(curr, ans)

		if prices[i] < bestbuy{
			bestbuy = prices[i]
		}
	}

	return ans
}
