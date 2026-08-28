/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
	nodeToNewNode := make(map [*Node]*Node)
	var dfs func(node *Node)*Node
	dfs = func(node *Node) *Node{
		if newNode, ok := nodeToNewNode[node]; ok{
			return newNode
		}
		newNode := &Node{
			Val: node.Val,
		}
		nodeToNewNode[node] = newNode

		for _, neighbor := range node.Neighbors{
			newNeighbor := dfs(neighbor)
			newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
		}
		return newNode
	}
	
	return dfs(node)
}
