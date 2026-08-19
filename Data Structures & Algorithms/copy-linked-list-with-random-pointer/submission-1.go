/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToCopy := make(map[*Node]*Node)
	oldToCopy[nil] = nil

	current := head
	for current != nil{
		if _, ok := oldToCopy[current]; !ok {
			oldToCopy[current] = &Node{Val:current.Val}
		}
		if current.Next != nil{
			if _, ok := oldToCopy[current.Next]; !ok{
				oldToCopy[current.Next] = &Node{Val: current.Next.Val}
			}
			oldToCopy[current].Next = oldToCopy[current.Next]
		}
		if current.Random != nil{
			if _, ok := oldToCopy[current.Random]; !ok{
				oldToCopy[current.Random] = &Node{Val: current.Random.Val}
			}
			oldToCopy[current].Random = oldToCopy[current.Random]
		}
		current = current.Next
	}
	return oldToCopy[head]
}
