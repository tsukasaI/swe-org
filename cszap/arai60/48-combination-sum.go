package arai60

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0)
	currentSum := 0
	var backtrack func(start int)
	backtrack = func(start int) {
		if currentSum > target {
			return
		}
		if currentSum == target {
			result = append(result, append(make([]int, 0, len(current)), current...))
			return
		}

		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i])
			currentSum += candidates[i]
			backtrack(i)
			currentSum -= candidates[i]
			current = current[:len(current)-1]
		}
	}
	backtrack(0)
	return result
}
