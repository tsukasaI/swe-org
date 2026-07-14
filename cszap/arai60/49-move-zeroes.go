package arai60

func moveZeroes(nums []int) {
	write := 0

	for write < len(nums) {
		if nums[write] == 0 {
			break
		}
		write++
	}

	for read := write; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
}
