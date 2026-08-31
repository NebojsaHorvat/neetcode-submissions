/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/* TODO optimize string space
// use littleEndian layour
15 - 1  -> Go back N nuber or nodes
15 - 0  -> ignore
14 - 1  -> new node is left copared to a parent
13 - 1  -> new node is right copared to a parent
12 - 1  -> nuber is negative and substrat 1000 from it
*/
import "encoding/binary"
type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	retBytes := []byte{}
	var goBackNodes uint16
	var dfs func(*TreeNode, bool)
	dfs = func(node *TreeNode, leftDir bool){
		if node == nil{
			return
		}
		if goBackNodes > 0 {
			val := goBackNodes
			val |= 0x8000
			serialize := make([]byte,2)
			binary.LittleEndian.PutUint16(serialize, val)
			retBytes = append(retBytes, serialize...)
			goBackNodes = 0
		}
		var val uint16
		val = uint16(node.Val)
		if leftDir{
			val |= 0x4000
		}else {
			val |= 0x2000
		}
		serialize := make([]byte,2)
		binary.LittleEndian.PutUint16(serialize, val)
		retBytes = append(retBytes, serialize...)
		dfs(node.Left, true)
		dfs(node.Right, false)
		goBackNodes++
	}
	dfs(root, true)
	ret := string(retBytes)
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	index := 0
	serialized := []byte(data)
	prevNodes := []*TreeNode{}
	var root *TreeNode
	for index < len(serialized){
		nodeValBytes := serialized[index:index+2]
		nodeVal := binary.LittleEndian.Uint16(nodeValBytes)
		index += 2
		comingFromLeft := false
		// Go back N nodes
		if nodeVal & 0x8000 > 0{
			nodesToReturn := int(nodeVal ^ 0x8000)
			// fmt.Println("Return ", nodesToReturn)
			prevNodes = prevNodes[:len(prevNodes) - nodesToReturn]
			continue
		}

		// Node is in the left side of the parent
		if nodeVal & 0x4000 > 0{ 
			nodeVal ^= 0x4000
			comingFromLeft = true
		// Node is in the right side of the parent
		}else if nodeVal & 0x2000 > 0{
			nodeVal ^= 0x2000
			comingFromLeft = false
		}else {
			panic("Node with wrong flag")
		}
		nodeValInt := int(nodeVal)
		newNode := &TreeNode{
			Val:nodeValInt,
		}
		// fmt.Println("NewNode ", nodeValInt)
		// Root
		if len( prevNodes ) == 0 {
			// fmt.Println("Adding root", nodeValInt)
			root = newNode
			prevNodes = append(prevNodes, newNode)
			continue
		}
		// Any other node
		parent := prevNodes[len(prevNodes)-1]
		if comingFromLeft{
			// fmt.Println("Parent: ", parent.Val," reciving left:", newNode.Val)
			parent.Left = newNode
		}else {
			// fmt.Println("Parent: ", parent.Val," reciving right:", newNode.Val)
			parent.Right = newNode
		}
		prevNodes = append(prevNodes, newNode)

	}
	return root
}
