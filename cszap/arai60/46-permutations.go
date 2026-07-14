package arai60

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			result = append(result, append(make([]int, 0, len(nums)), current...))
			return
		}
		for i := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			current = append(current, nums[i])
			backtrack()
			used[i] = false
			current = current[:len(current)-1]
		}
	}

	backtrack()
	return result
}
