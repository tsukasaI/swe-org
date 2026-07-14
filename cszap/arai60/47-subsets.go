package arai60

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0, len(nums))

	var backtrack func(start int)
	backtrack = func(start int) {
		result = append(result, append(make([]int, 0, len(current)), current...))
		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}
	backtrack(0)
	return result
}
