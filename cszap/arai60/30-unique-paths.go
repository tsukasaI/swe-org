package arai60

func uniquePaths(m int, n int) int {
	dp := make([]int, m)
	for i := range m {
		dp[i] = 1
	}

	for range n {
		for i := range m {
			if i == 0 {
				dp[0] = 1
			} else {
				dp[i] = dp[i-1] + dp[i]
			}
		}
	}
	return dp[m-1]
}
