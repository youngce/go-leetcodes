package solutions

import "math"

func dfsHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := dfsHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := dfsHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1

}

func isBalanced(root *TreeNode) bool {

	return dfsHeight(root) != -1
}
