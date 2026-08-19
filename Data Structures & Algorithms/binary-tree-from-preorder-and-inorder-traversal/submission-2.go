/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	ret, _ := buildTreeRec(preorder, inorder)
	return ret
}

func buildTreeRec(preorder []int, inorder []int) (*TreeNode, []int) {
	if len(inorder) == 0  || len(preorder) == 0{
		return nil, preorder
	}

	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}
	preorder = preorder[1:]

	// find left partition
	i := 0
	for ; i<len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	leftInorder := inorder[:i]
	var rightInorder []int
	if i+1 >= len(inorder){
		rightInorder = nil
	}else{
		rightInorder = inorder[i+1:]
	}
	// fmt.Println("Root: ", rootVal)
	// fmt.Println("LeftInorder:", leftInorder)
	// fmt.Println("RightInorder:", rightInorder)
	root.Left, preorder = buildTreeRec(preorder, leftInorder)
	root.Right, preorder = buildTreeRec(preorder, rightInorder)
	return root, preorder
}