/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
		return false
	}
	if root.Val == subRoot.Val && compareTrees(root, subRoot){
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func compareTrees(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil{
		return true
	} else if root == nil || subRoot == nil{
		return false
	}
	if root.Val != subRoot.Val{
		return false
	}
	return compareTrees(root.Left, subRoot.Left) && compareTrees(root.Right, subRoot.Right)
}
