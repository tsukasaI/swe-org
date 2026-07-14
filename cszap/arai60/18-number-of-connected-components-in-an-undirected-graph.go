package arai60

func countEdgeSum(edges [][]int, n int) int {
	adjacentList := make(map[int][]int)
	for _, v := range edges {
		adjacentList[v[0]] = append(adjacentList[v[0]], v[1])
		adjacentList[v[1]] = append(adjacentList[v[1]], v[0])
	}

	visited := make(map[int]struct{})
	var dfs func(i int)
	dfs = func(i int) {
		if _, ok := visited[i]; !ok {
			visited[i] = struct{}{}
			for _, j := range adjacentList[i] {
				dfs(j)
			}
		}
	}

	count := 0
	for i := range n {
		if _, ok := visited[i]; !ok {
			count++
			dfs(i)
		}
	}
	return count
}
