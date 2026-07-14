package arai60

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := -10000
	for _, v := range nums {
		sum += v
		maxSum = max(sum, maxSum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
