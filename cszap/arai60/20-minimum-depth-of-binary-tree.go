package arai60

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// check if the root is a leaf node
	if root.Left == nil && root.Right == nil {
		// leaf node
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
