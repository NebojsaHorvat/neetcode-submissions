type MaxHeap struct{
	data []int
}

func (h *MaxHeap) Len ()int { return len(h.data)}
func (h *MaxHeap) Push(val int){
	h.data = append(h.data, val)
	h.siftUp(h.Len()-1)
}

func(h *MaxHeap) Pop() int{
	top := h.data[0]
	h.data[0] = h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	if h.Len() > 0 {
		h.siftDown(0)
	}
	return top
}

func Constructor (data []int) MaxHeap{
	maxHeap := MaxHeap{}
	for _, val := range data {
		maxHeap.Push(val)
	}
	return maxHeap
}

func (h *MaxHeap) siftUp (i int){
	for i > 0 {
		parent := (i-1)/2
		if h.data[parent] >= h.data[i]{
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MaxHeap) siftDown (i int){
	n := len(h.data)
	largest := i
	for {
		left := i*2 +1
		right := i*2 +2
		if left < n && h.data[left] > h.data[largest]{
			largest = left
		}
		if right < n && h.data[right] > h.data[largest]{
			largest = right
		}
		if largest == i{
			break
		}
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

func lastStoneWeight(stones []int) int {
	maxHeap := Constructor(stones)
	for maxHeap.Len() > 1{
		x := maxHeap.Pop()
		y := maxHeap.Pop()
		if x == y{
			continue
		}
		remain := 0
		if x>y{
			remain = x-y
		}else{
			remain = y-x
		}
		maxHeap.Push(remain)
	}
	if maxHeap.Len() == 1{
		return maxHeap.Pop()
	}else{
		return 0
	}
}
