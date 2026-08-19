/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	reverted := -1
	ret := -1
	reverseBack(root, k, &reverted, &ret)

	return ret

}

func reverseBack(root *TreeNode, k int, reversed *int, ret *int){
	if root == nil {
		if *reversed == -1{
			*reversed = 1
		}
		return
	}
	reverseBack(root.Left, k, reversed, ret)
	// fmt.Println("Current root:",root.Val, " Current reversed:",*reversed)
	if k == *reversed{
		*ret = root.Val
	}
	*reversed++
	if root.Right != nil{	
		reverseBack(root.Right,k,reversed, ret)
		// fmt.Println("Current root:",root.Val, " Current reversed:",*reversed)
		if k == *reversed{
			*ret = root.Val
		}
		// *reversed++
	}
	
}
