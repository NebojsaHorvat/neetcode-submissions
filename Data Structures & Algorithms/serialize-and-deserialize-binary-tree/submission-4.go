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
	strSlice := []string{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode){
		if node == nil{
			strSlice = append(strSlice, "N")
			return
		}
		strVal := strconv.Itoa(node.Val)
		strSlice = append(strSlice, strVal)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	ret := strings.Join(strSlice,",")
	fmt.Println(ret)
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data,",")
	index := 0
	var dfs func([]string) *TreeNode
	dfs = func(nodes []string)*TreeNode{
		if index >= len(nodes){
			return nil
		}
		nodeVal := nodes[index]
		if nodeVal == "N"{
			return nil
		}
		nodeValInt, _ := strconv.Atoi(nodeVal)
		node := TreeNode{
			Val:nodeValInt,
		}
		index++
		node.Left = dfs(nodes)
		index++
		node.Right = dfs(nodes)
		return &node
	}
	return dfs(nodes)
}
