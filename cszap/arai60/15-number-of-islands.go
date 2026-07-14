package arai60

func numIslands(grid [][]byte) int {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}
