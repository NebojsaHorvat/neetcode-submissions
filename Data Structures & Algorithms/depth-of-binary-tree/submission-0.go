/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    maxDepth := 0
	calculateDepth(root, 0, &maxDepth)
	return maxDepth
}

func calculateDepth(root *TreeNode, depth int, maxDepth *int){
	if root == nil {
		if depth > *maxDepth{
			*maxDepth = depth
		}
		return
	}
	calculateDepth(root.Left, depth+1, maxDepth)
	calculateDepth(root.Right, depth+1, maxDepth)
}