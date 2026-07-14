package arai60

func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, upper, lower int) bool
	dfs = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}

		if root.Val <= lower || root.Val >= upper {
			return false
		}

		return dfs(root.Left, lower, root.Val) && dfs(root.Right, root.Val, upper)
	}

	return dfs(root, -2<<31, 2<<31)
}
