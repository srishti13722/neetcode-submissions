/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil{
		return
	}
    //find mid
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	//reverse second half
	
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil

	for curr != nil{
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	// add nodes
	h1 := head
	h2 := prev
	
	for h2 != nil{
		n1 := h1.Next
		n2 := h2.Next

		h1.Next = h2
		h2.Next = n1

		h1 = n1
		h2 = n2
	}

}
