package arai60

func splitBST(root *Node, target int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}

	if root.Val <= target {
		smaller, larger := splitBST(root.Right, target)
		root.Right = smaller
		return root, larger
	} else {
		smaller, larger := splitBST(root.Left, target)
		root.Left = larger
		return smaller, root
	}
}
