package arai60

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		element := make([]int, 0, len(queue))
		nextQueue := make([]*TreeNode, 0, len(queue)*2)
		for _, node := range queue {
			element = append(element, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		result = append(result, element)
		queue = nextQueue
	}
	return result
}
