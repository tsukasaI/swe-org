// https://leetcode.com/problems/subarray-sum-equals-k/description/
package arai60

func subarraySum(nums []int, k int) int {
	sum := 0
	result := 0
	hm := make(map[int]int)
	hm[0] = 1
	for i := range nums {
		sum += nums[i]
		if _, ok := hm[sum-k]; ok {
			result += hm[sum-k]
		}
		hm[sum]++
	}
	return result
}
