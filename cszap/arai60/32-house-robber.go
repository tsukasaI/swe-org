package arai60

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l, r := nums[0], max(nums[0], nums[1])

	result := r
	for i := 2; i < len(nums); i++ {
		result = max(l+nums[i], r)
		l = r
		r = result
	}
	return result
}
