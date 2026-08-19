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
			if q[i].Left != nil{
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil{
				q = append(q, q[i].Right)
			}
		}
		q = q[lenQ:]
	}
	return ret
}
