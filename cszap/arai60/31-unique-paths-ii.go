package arai60

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	h, v := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([]int, v)
	for i, e := range obstacleGrid[0] {
		if e == 1 {
			dp[i] = 0
			break
		} else {
			dp[i] = 1
		}
	}

	for i := 1; i < h; i++ {
		for j := range v {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j == 0 {
				dp[j] = dp[j]
			} else {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}
	return dp[v-1]
}
