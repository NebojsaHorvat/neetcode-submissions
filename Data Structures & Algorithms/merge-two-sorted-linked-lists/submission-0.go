/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head, current *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil{
		return list1
	}
	if list1.Val < list2.Val{
		current = list1
		head = list1
		list1 = list1.Next
	}else{
		current = list2
		head = list2
		list2 = list2.Next
	}
	
	for list1 != nil || list2 != nil {
		if list1 == nil{
			// fmt.Println("Current: ",current)
			current.Next = list2
			// fmt.Println("exiting because list1 is nil")
			return head
		}
		if list2 == nil{
			// fmt.Println("Current: ",current)
			current.Next = list1
			// fmt.Println("exiting because list2 is nil")
			return head
		}
		if list1.Val < list2.Val{
			current.Next = list1
			// fmt.Println("Current.Next: ",list1)
			list1 = list1.Next
			// fmt.Println("List1: ",list1,"\n***")
		}else{
			current.Next = list2
			// fmt.Println("Current.Next: ",list2)
			list2 = list2.Next
			// fmt.Println("List2: ",list2,"\n***")
		}
		current = current.Next
	}
	return head
}
