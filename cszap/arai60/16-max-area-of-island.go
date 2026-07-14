package arai60

func maxAreaOfIsland(grid [][]int) int {
	area := 0

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}

	for i := range grid {
		for j := range grid[0] {
			area = max(dfs(i, j), area)
		}
	}
	return area
}
