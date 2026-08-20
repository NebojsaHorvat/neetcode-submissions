/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO optimize string space
// bytes layout
// val1 val2 left1 left2 right1 right2...
// bit layout of val1			bit layout of val2
// exists|sign| val |...		|val| val | val

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ret []string 

    var dfs func ( *TreeNode)
	dfs = func (root *TreeNode){
		if root == nil{
			ret = append(ret, "N")
			return
		}
		// fmt.Println("Root: ",strconv.Itoa(root.Val))
		ret = append(ret, strconv.Itoa(root.Val))
		// fmt.Println(ret)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	// var retString string
	// for _, retS := range ret{
	// 	retString = retString+ retS + ","
	// }
	retString := strings.Join(ret,",")
	fmt.Println(retString)
	return retString
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strNodes := strings.Split(data,",")
	fmt.Println(strNodes)
	var dfs func()*TreeNode
	dfs = func()*TreeNode{
		if len(strNodes) == 0 {
			return nil
		}
		strNode := strNodes[0]
		strNodes = strNodes[1:]
		// fmt.Println(strNodes)
		if strNode == "N"{
			return nil
		}

		val,_ := strconv.Atoi(strNode)
		newNode := &TreeNode{
			Val:val,
		}
		newNode.Left = dfs()
		newNode.Right = dfs()

		return newNode
	}
	return dfs()

    return nil
}
