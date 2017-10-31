package solutions

import (
	"math"
	"debug/dwarf"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func isSymmetric(root *TreeNode) bool {

	return root==nil||isSymmetricHelp(root.Left,root.Right)

}
func isSymmetricHelp(left *TreeNode, right *TreeNode) bool{
	if left==nil||right==nil{
		return left==right
	}
	if left.Val!=right.Val{
		return false
	}
	return isSymmetricHelp(left.Left,right.Right)&&
		isSymmetricHelp(left.Right,right.Left)
}