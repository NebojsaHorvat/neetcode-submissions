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
	retStrs := []string{}
	q := []*TreeNode{root}
	
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil{
			retStrs = append(retStrs, "N")
			continue
		}else{
			retStrs = append(retStrs, strconv.Itoa(node.Val))
		}
		q = append(q, node.Left)
		q = append(q, node.Right)
	}
	ret := strings.Join(retStrs,",")
	fmt.Println(ret)
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strNodes := strings.Split(data,",")
	if strNodes[0] == "N"{
		return nil
	}
	val, _ := strconv.Atoi(strNodes[0])
	root := &TreeNode{Val: val,}
	q := []*TreeNode{root}
	// Cover left node if true, Cover right if false
	leftNode:= true
	for _, strNode := range strNodes[1:]{
		var node *TreeNode 
		if strNode != "N"{
			val, _ := strconv.Atoi(strNode)
			node = &TreeNode{Val: val,}
		}
		if leftNode{ 
			if q[0]!= nil{
				q[0].Left = node
			}
			leftNode = false
			if node != nil{
				q = append(q, node)
			}
		}else{
			if q[0]!= nil{
				q[0].Right = node
			}
			leftNode = true
			if node != nil{
				q = append(q, node)
			}
			q = q[1:]
		}
	}
	return root
}
