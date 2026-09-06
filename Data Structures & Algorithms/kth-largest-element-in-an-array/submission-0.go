type MinHeap struct{
	data []int
	k int
}

func (h *MinHeap) Len ()int { return len(h.data)}
func (h *MinHeap) Push(val int){
	if h.Len() < h.k {
		h.data = append(h.data, val)
		h.siftUp(h.Len()-1)
		return
	}
	if val > h.data[0]{
		h.data[0] = val
		h.siftDown(0)
	}
	
}

func(h *MinHeap) Pop() int{
	if h.Len() == 0{
		return 0
	}
	top := h.data[0]
	h.data[0] = h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	if h.Len() > 0 {
		h.siftDown(0)
	}
	return top
}

func (h *MinHeap) siftUp (i int){
	for i > 0 {
		parent := (i-1)/2
		if h.data[parent] <= h.data[i]{
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MinHeap) siftDown (i int){
	n := len(h.data)
	smallest := i
	for {
		left := i*2 +1
		right := i*2 +2
		if left < n && h.data[left] < h.data[smallest]{
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest]{
			smallest = right
		}
		if smallest == i{
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func findKthLargest(nums []int, k int) int {
	minHeap := MinHeap{[]int{}, k}
	for _, num := range nums{
		minHeap.Push(num)
	}
	return minHeap.Pop()
}
