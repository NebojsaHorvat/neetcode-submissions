/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
    current := head
	var prev *ListNode = nil
	for current.Next != nil{
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	current.Next = prev
	return current
}
