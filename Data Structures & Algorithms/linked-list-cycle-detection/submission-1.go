/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    fast := head; slow := head

	for fast != nil {
		fast = fast.Next
		if fast == slow{
			return true
		}
		if fast == nil{
			return false
		}
		fast = fast.Next
		slow = slow.Next
		if fast == slow{
			return true
		}
	}
	return false
}
