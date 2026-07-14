package arai60

import (
	"sort"
)

func lengthoflis(nums []int) int {
	tails := make([]int, 0)
	tails = append(tails, nums[0])

	for i := 1; i < len(nums); i++ {
		if tails[len(tails)-1] < nums[i] {
			tails = append(tails, nums[i])
		} else {
			index := sort.SearchInts(tails, nums[i])
			tails[index] = nums[i]
		}
	}
	return len(tails)
}
