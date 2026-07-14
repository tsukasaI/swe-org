package arai60

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		if c := (l + r) / 2; nums[c] >= target {
			r = c
		} else {
			l = c + 1
		}
	}
	return l
}
