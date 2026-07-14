package arai60

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	subTargetSum := targetSum - root.Val
	return hasPathSum(root.Left, subTargetSum) || hasPathSum(root.Right, subTargetSum)
}
