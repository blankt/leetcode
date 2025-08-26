package tree

func findKthLargest(nums []int, k int) int {
	heap := NewHeap[ComparableInt](k)
	for _, num := range nums {
		heap.Push(ComparableInt(num))
	}

	result, ok := heap.Top()
	if !ok {
		return -1
	}
	return int(result)
}

type Comparable[T any] interface {
	Less(other T) bool
}

type ComparableInt int

func (a ComparableInt) Less(b ComparableInt) bool {
	return a <= b
}

type Heap[T Comparable[T]] struct {
	data []T
	size int
}

func NewHeap[T Comparable[T]](size int) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0, size),
		size: size,
	}
}

func (h *Heap[T]) Push(x T) {
	if len(h.data) < h.size {
		h.data = append(h.data, x)
		h.pop()
	} else {
		if !h.data[0].Less(x) {
			return
		} else {
			h.data[0] = x
			h.down()
		}
	}
}

func (h *Heap[T]) Top() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	return h.data[0], true
}

func (h *Heap[T]) pop() {
	i := len(h.data) - 1
	for {
		parent := (i - 1) / 2
		if h.data[parent].Less(h.data[i]) {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *Heap[T]) down() {
	i := 0
	for {
		var son int
		left := 2*i + 1
		right := 2*i + 2
		if left >= len(h.data) {
			break
		}
		if right < len(h.data) && h.data[right].Less(h.data[left]) {
			son = right
		} else {
			son = left
		}

		if h.data[i].Less(h.data[son]) {
			break
		}

		h.data[i], h.data[son] = h.data[son], h.data[i]
		i = son
	}
}

func (h *Heap[T]) Pop() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	top := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down()
	return top, true
}
