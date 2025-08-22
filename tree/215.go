package tree

func findKthLargest(nums []int, k int) int {
	heap := NewHeap(func(a, b int) bool {
		return a <= b
	}, k)
	for _, num := range nums {
		heap.Push(num)
	}
	return heap.Top()
}

type Heap struct {
	data []int
	less func(int, int) bool
	size int
}

func NewHeap(less func(int, int) bool, size int) *Heap {
	return &Heap{
		data: make([]int, 0, size),
		less: less,
		size: size,
	}
}

func (h *Heap) Push(x int) {
	if len(h.data) < h.size {
		h.data = append(h.data, x)
		h.pop()
	} else {
		if h.data[0] >= x {
			return
		} else {
			h.data[0] = x
			h.down()
		}
	}
}

func (h *Heap) Top() int {
	if len(h.data) == 0 {
		return 0
	}
	return h.data[0]
}

func (h *Heap) pop() {
	i := len(h.data) - 1
	for {
		parent := (i - 1) / 2
		if h.less(h.data[parent], h.data[i]) {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *Heap) down() {
	i := 0
	for {
		var son int
		left := 2*i + 1
		right := 2*i + 2
		if left >= len(h.data) {
			break
		}
		if right < len(h.data) && h.less(h.data[right], h.data[left]) {
			son = right
		} else {
			son = left
		}

		if h.less(h.data[i], h.data[son]) {
			break
		}

		h.data[i], h.data[son] = h.data[son], h.data[i]
		i = son
	}
}
