package arai60

import (
	"slices"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 0 {
		return nil
	}

	nodeVal := preorder[0]
	nodeValIndexInInorder := slices.Index(inorder, nodeVal)

	return &TreeNode{
		Val:   nodeVal,
		Left:  buildTree(preorder[1:1+nodeValIndexInInorder], inorder[:nodeValIndexInInorder]),
		Right: buildTree(preorder[1+nodeValIndexInInorder:], inorder[nodeValIndexInInorder+1:]),
	}
}
