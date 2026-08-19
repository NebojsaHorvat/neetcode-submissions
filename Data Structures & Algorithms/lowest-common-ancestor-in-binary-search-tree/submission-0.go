/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    traversedMap := make(map [*TreeNode]bool)
	var lowestNode *TreeNode
	findElementP(root,p,traversedMap)
	findElementQ(root,q,traversedMap,&lowestNode)

	return lowestNode
}

func findElementP(root *TreeNode, e *TreeNode, traversedMap map[*TreeNode]bool){
	traversedMap[root] = true
	if root == e {
		return 
	}
	if e.Val < root.Val {
		findElementP(root.Left, e, traversedMap)
	}else{
		findElementP(root.Right, e, traversedMap)
	}
}

func findElementQ(root *TreeNode, e *TreeNode, traversedMap map[*TreeNode]bool, lowestNode **TreeNode){
	fmt.Println("FindingQ, element:",root.Val)
	if _, ok := traversedMap[root]; ok {
		fmt.Println("Found element in P:", root.Val)
		*lowestNode = root
	}else {
		fmt.Println("Not found element in P:", root.Val)
		return
	}
	
	if root == e{
		return
	}

	if e.Val < root.Val {
		findElementQ(root.Left, e, traversedMap, lowestNode)
	}else{
		findElementQ(root.Right, e, traversedMap, lowestNode)
	}
}
