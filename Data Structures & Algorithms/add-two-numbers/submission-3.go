/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := l1
	res := 0
	prev := l1
    for l1 != nil && l2!= nil{
		sum := l1.Val + l2.Val
		if res > 0{
			sum += res
			res = 0
		}
		if sum > 9{
			sum-=10
			res = 1
		}
		l1.Val = sum

		prev = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if res > 0 && l2 == nil && l1 == nil{
		prev.Next = &ListNode{Val:1}
		return ret
	}

	if l1 == nil && l2 != nil{
		fmt.Println("Converting l2 to l1")
		l1 = l2
		l2 = nil
		prev.Next = l1
	}
	if res == 0 {
		return ret
	}

	if l2 == nil && l1 != nil && res > 0{
		for l1 != nil{
			if res > 0{
				l1.Val += res
				res = 0
			}
			if l1.Val > 9{
				l1.Val-=10
				res = 1
			}
			prev = l1
			l1 = l1.Next
		}
		if res > 0{
			prev.Next = &ListNode{Val:1}
			return ret
		}
	}
	

	return ret
}
