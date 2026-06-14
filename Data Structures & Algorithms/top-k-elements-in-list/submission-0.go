type Pair struct{
	freq int
	num int
}

type MinHeap []Pair

func(h MinHeap)Len()int{return len(h)}
func(h MinHeap)Less(i , j int)bool{return h[i].freq < h[j].freq}
func(h MinHeap)Swap(i , j int){h[i], h[j] = h[j], h[i]}
func(h *MinHeap)Push(data any){
	*h  = append(*h , data.(Pair))
}

func(h *MinHeap)Pop()any{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)

	for _ , num := range nums{
		mp[num]++
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for num , freq := range mp{
		heap.Push(minHeap, Pair{freq, num})

		if minHeap.Len() > k{
			heap.Pop(minHeap)
		}
	}

	var res []int

	for minHeap.Len()>0{
		top := heap.Pop(minHeap).(Pair)
		ele := top.num
		res = append(res, ele)
	}

	return res
}
