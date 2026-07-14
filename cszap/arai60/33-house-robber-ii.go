package arai60

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robLinier(nums[:len(nums)-1]), robLinier(nums[1:]))
}

func robLinier(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, 0, len(nums))

	dp = append(dp, nums[0])
	dp = append(dp, max(nums[0], nums[1]))

	for i := 2; i < len(nums); i++ {
		dp = append(dp, max(dp[i-2]+nums[i], dp[i-1]))
	}
	return dp[len(dp)-1]
}
