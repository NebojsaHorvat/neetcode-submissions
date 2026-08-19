/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    ret := []int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		lenQ := len(q)
		seenOnThisLvl := q[lenQ-1].Val
		ret = append(ret, seenOnThisLvl)
		for i:=0; i<lenQ; i++ {
			element := q[0]
			q = q[1:]
			if element.Left != nil{
				q = append(q, element.Left)
			}
			if element.Right != nil{
				q = append(q, element.Right)
			}
		}
	}
	return ret
}
