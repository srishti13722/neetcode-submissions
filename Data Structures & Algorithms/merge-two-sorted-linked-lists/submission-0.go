/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head1 := list1
	head2 := list2
	
	var d *ListNode = &ListNode{Val : -1}
	var dummy *ListNode = d

	for head1 != nil || head2 != nil{
		if head1 != nil && head2 != nil{
			if head1.Val <= head2.Val{
				dummy.Next = &ListNode{Val : head1.Val}
				head1 = head1.Next
			}else{
				dummy.Next = &ListNode{Val : head2.Val}
				head2 = head2.Next
			}
		}else if head1 != nil{
			dummy.Next = &ListNode{Val : head1.Val}
			head1 = head1.Next
		}else{
			dummy.Next = &ListNode{Val : head2.Val}
			head2 = head2.Next
		}

		dummy = dummy.Next
	}

	return d.Next
}
