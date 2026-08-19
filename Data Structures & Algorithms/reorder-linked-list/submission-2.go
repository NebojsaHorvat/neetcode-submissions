/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head.Next == nil  || head.Next.Next == nil{
		return
	}
	slow := head
	fast := head
	var slowPrev *ListNode
	evenNumberOfElements := true
	for fast != nil{
		fast = fast.Next
		if fast == nil{
			fmt.Println("breaking in half")
			evenNumberOfElements = false
			break
		}
		fast = fast.Next
		slowPrev = slow
		slow = slow.Next
	}
	if evenNumberOfElements{
		slowPrev.Next = nil
	}
	fmt.Println("Half",slow)
	// slow is in the middle of the list
	// revesle the list after middle
	current := slow
	var prev *ListNode
	for current != nil{
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	// prev is the last element
	// now create the required list
	backHead := prev
	frontHead := head
	fmt.Println("BackHead:",backHead, "  BackHead +1 : ",backHead.Next, "  BackHead+2 ", backHead.Next.Next)
	fmt.Println("FrontHead:",frontHead, "  FrontHead +1 : ",frontHead.Next, "  FrontHead+2 ", frontHead.Next.Next)

	for backHead != nil && frontHead != nil{
		// Save next front element
		frontNext := frontHead.Next
		// Put back element to front
		frontHead.Next = backHead
		// Move back head
		backNext := backHead.Next
		// Set Back element to next Front
		backHead.Next = frontNext

		// MoveFront and back
		frontHead = frontNext
		backHead = backNext
	}
}	
