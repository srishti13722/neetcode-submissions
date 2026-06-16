type MaxHeap []int
func(h MaxHeap)Len()int{return len(h)}
func(h MaxHeap)Less(i, j int)bool{return h[i] > h[j]}
func(h MaxHeap)Swap(i, j int){h[i], h[j] = h[j], h[i]}
func(h *MaxHeap)Push(data any){*h = append(*h, data.(int))}
func(h *MaxHeap)Pop()any{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func(h MaxHeap)Top()int{
	return h[0]
}

func NewMaxHeap()*MaxHeap{
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	return maxHeap
}

type MinHeap []int
func(h MinHeap)Len()int{return len(h)}
func(h MinHeap)Less(i, j int)bool{return h[i] < h[j]}
func(h MinHeap)Swap(i, j int){h[i], h[j] = h[j], h[i]}
func(h *MinHeap)Push(data any){*h = append(*h, data.(int))}
func(h *MinHeap)Pop()any{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func(h MinHeap)Top()int{
	return h[0]
}

func NewMinHeap()*MinHeap{
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	return minHeap
}

type MedianFinder struct {
    maxHeap *MaxHeap
	minHeap *MinHeap
}


func Constructor() MedianFinder {
	maxHeap := NewMaxHeap()
	minHeap := NewMinHeap()
    return MedianFinder{
		maxHeap : maxHeap,
		minHeap : minHeap,
	}
}


func (this *MedianFinder) AddNum(num int)  {
    if this.maxHeap.Len() == 0{
		heap.Push(this.maxHeap, num)
	}else{
		if num > this.maxHeap.Top(){
			heap.Push(this.minHeap, num)
			if this.minHeap.Len() > this.maxHeap.Len(){
				heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
			}
		}else{
			heap.Push(this.maxHeap, num)
			if this.maxHeap.Len() > this.minHeap.Len()+1{
				heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
			}
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
    if this.minHeap.Len() == this.maxHeap.Len(){
		return float64(this.minHeap.Top() + this.maxHeap.Top()) / 2
	}

	return float64(this.maxHeap.Top())
}