package arai60

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if m := (l + r) / 2; nums[m] == target {
			return m
		} else if nums[l] <= nums[m] {
			if nums[l] <= target && nums[m] >= target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] <= target && nums[r] >= target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
