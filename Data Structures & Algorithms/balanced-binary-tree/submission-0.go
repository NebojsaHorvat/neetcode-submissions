/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    _, ret := reverseHeight(root)
	return !ret
}

func reverseHeight(root *TreeNode) (int, bool){
	if root == nil{
		return 0, false
	}
	leftHeight, foundProblemRight := reverseHeight(root.Left)
	leftHeight++
	rightHeight, foundProblemLeft := reverseHeight(root.Right)
	rightHeight++
	if foundProblemLeft || foundProblemRight {
		return 0, true
	}
	if leftHeight - rightHeight > 1  || rightHeight - leftHeight > 1{
		return 0, true
	}
	if leftHeight > rightHeight{
		return leftHeight, false
	}else {
		return rightHeight, false
	}
}
