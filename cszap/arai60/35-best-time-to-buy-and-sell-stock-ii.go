package arai60

func maxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
