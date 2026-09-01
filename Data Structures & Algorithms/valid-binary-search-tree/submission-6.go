/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return DFS(root, 10000000, -10000000)
}

func DFS(node *TreeNode, left, right int) bool {
	if node == nil{
		return true
	}
	if node.Val >= left || node.Val <= right{
		return false
	}
	return DFS(node.Left, min(node.Val, left), right) && DFS(node.Right, left, max(right,node.Val))
}