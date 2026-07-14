package arai60

func minSubArrayLen(target int, nums []int) int {
	l := 0
	sum := 0
	const initLen = 1000000
	minLen := initLen
	for r := 0; r < len(nums); r++ {
		sum += nums[r]

		for sum >= target {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if minLen == initLen {
		return 0
	}
	return minLen
}
