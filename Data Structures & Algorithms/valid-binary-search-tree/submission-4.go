/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return DFS(root, math.MinInt64, math.MaxInt64)
}

func DFS(root *TreeNode, left, right int) bool {
	if root == nil{
		return true
	}
	if root.Val <= left || root.Val >= right{
		return false
	}
	return DFS(root.Left, left, root.Val) && DFS (root.Right, root.Val, right)
}