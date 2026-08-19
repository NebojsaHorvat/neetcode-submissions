/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return DFS(root, root.Val)
}

func DFS (root *TreeNode, maxVal int) int{
	if root == nil{
		return 0
	}
	isCurrentNodeOk := 0
	if root.Val >= maxVal{
		isCurrentNodeOk++
	}
	
	if root.Val > maxVal {
		maxVal = root.Val
	}
	return isCurrentNodeOk+ DFS(root.Left, maxVal) + DFS(root.Right, maxVal)
}
