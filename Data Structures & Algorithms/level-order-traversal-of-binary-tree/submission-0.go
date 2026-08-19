/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int,0)
	q := []*TreeNode{root}
    if root == nil {
		return ret
	}
	for len(q) > 0 {
		lenQ := len(q)
		level := make([]int,0)
		for i:=0 ; i<lenQ; i++ {
			element := q[0]
			q = q[1:]
			level = append(level, element.Val)
			if element.Left != nil{
				q = append(q, element.Left)
			}
			if element.Right != nil{
				q = append(q, element.Right)
			}
		}
		ret = append(ret, level)
	}
	return ret
}
