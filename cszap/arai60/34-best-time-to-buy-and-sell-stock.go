package arai60

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))

	minPrice := prices[0]

	for i := 1; i < len(dp); i++ {
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return dp[len(dp)-1]
}
