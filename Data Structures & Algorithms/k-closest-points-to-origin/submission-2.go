type MinHeap struct{
	data [][]int
}

func (h *MinHeap) Len ()int { return len(h.data)}
func (h *MinHeap) Push(val []int){
	h.data = append(h.data, val)
	h.siftUp(h.Len()-1)
}

func(h *MinHeap) Pop() []int{
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
		if h.data[parent][2] <= h.data[i][2]{
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
		if left < n && h.data[left][2] < h.data[smallest][2]{
			smallest = left
		}
		if right < n && h.data[right][2] < h.data[smallest][2]{
			smallest = right
		}
		if smallest == i{
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func Constructor (data [][]int) MinHeap{
	minHeap := MinHeap{}
	for _, val := range data {
		minHeap.Push(val)
	}
	return minHeap
}

func kClosest(points [][]int, k int) [][]int {
	minHeap := Constructor([][]int{})
	for _, point := range points{
		x := point[0]
		y := point[1]
		val := []int{x,y, x*x + y*y}
		minHeap.Push(val)
	}
	ret := [][]int{}
	for i:=0; i<k; i++{
		val := minHeap.Pop()
		ret = append(ret, []int{val[0],val[1]})
	}
	return ret
}
