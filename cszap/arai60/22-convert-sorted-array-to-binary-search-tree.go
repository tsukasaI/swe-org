package arai60

func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}

	center := size / 2

	return &TreeNode{
		Val:   nums[center],
		Left:  sortedArrayToBST(nums[:center]),
		Right: sortedArrayToBST(nums[center+1:]),
	}
}
