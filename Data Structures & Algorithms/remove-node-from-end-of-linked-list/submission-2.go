/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    num := 0
	current := head
	for current != nil{
		current = current.Next
		num++
	}
	numToRemove := num-n-1
	if numToRemove < 0 {
		return head.Next
	}
	current = head
	for numToRemove > 0{
		current = current.Next
		numToRemove--
	}
	fmt.Println(current)
	if num == 1{
		return nil
	}
	current.Next = current.Next.Next
	return head
}
