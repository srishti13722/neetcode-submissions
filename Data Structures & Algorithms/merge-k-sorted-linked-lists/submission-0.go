/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type MinHeap []*ListNode

func(h MinHeap)Len()int{return len(h)}
func(h MinHeap)Less(i, j int)bool{return h[i].Val < h[j].Val}
func(h MinHeap)Swap(i, j int){h[i], h[j] = h[j], h[i]}
func(h *MinHeap)Push(data any){*h = append(*h, data.(*ListNode))}
func(h *MinHeap)Pop()any{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, head := range lists{
		heap.Push(minHeap, head)
	}

	dummy := &ListNode{}
	curr := dummy
	for minHeap.Len() != 0{
		top := heap.Pop(minHeap).(*ListNode)
		curr.Next = top
		curr = curr.Next
		if top.Next != nil{
			heap.Push(minHeap, top.Next)
		}		
	}

	return dummy.Next
}
