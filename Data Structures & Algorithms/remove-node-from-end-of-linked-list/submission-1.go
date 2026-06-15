/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {    
	fast := head

	for n != 0{
		fast = fast.Next
		n = n-1
	}

	if fast == nil{
		return head.Next
	}

	slow := head

	for fast != nil && fast.Next!=nil{
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
	
}
