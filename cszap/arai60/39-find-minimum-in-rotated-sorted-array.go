package arai60

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		if m := (l + r) / 2; nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
