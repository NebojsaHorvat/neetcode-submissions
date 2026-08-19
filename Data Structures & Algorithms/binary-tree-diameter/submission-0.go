/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := maxHeight(root.Left)
	rightHeight := maxHeight(root.Right)
	diameter := leftHeight + rightHeight

	subTreeDiameter := max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))

	return max(diameter, subTreeDiameter)
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxHeight(root.Right), maxHeight(root.Left))
}

func max (a, b int) int{
	if a>b {
		return a
	}
	return b
}