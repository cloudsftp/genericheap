package genericheap

// ComparisonFunc defines the order in which elements are popped off the heap
type ComparisonFunc[T any] func(a, b T) int

type Heap[T any] struct {
	data []T
	cmp  ComparisonFunc[T]
}

// NewHeap creates a new generic heap
// It is a min heap that orderes its elements according to the cmp function
//
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b or a and b are incomparable in the sense of
// a strict weak ordering.
func NewHeap[T any](cmp ComparisonFunc[T]) Heap[T] {
	return Heap[T]{
		data: []T{},
		cmp:  cmp,
	}
}

// Push a new value to the heap
func (h *Heap[T]) Push(value T) {
	index := len(h.data)
	h.data = append(h.data, value)

	for index > 0 {
		nextIndex := (index - 1) / 2
		if h.less(nextIndex, index) {
			break
		}

		h.swap(index, nextIndex)
		index = nextIndex
	}
}

// Pop the minimal value from the heap
func (h *Heap[T]) Pop() (*T, bool) {
	if len(h.data) == 0 {
		return nil, false
	}

	value := h.data[0]

	// move last element to root
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	index := 0
	for {
		indexLeft := 2*index + 1
		indexRight := indexLeft + 1

		nextIndex := indexLeft
		if h.less(indexRight, indexLeft) {
			nextIndex = indexRight
		}

		if nextIndex >= len(h.data) ||
			h.less(index, nextIndex) {
			break
		}

		h.swap(index, nextIndex)
		index = nextIndex
	}

	return &value, true
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) less(i, j int) bool {
	if i >= len(h.data) || j >= len(h.data) {
		return false
	}

	return h.cmp(h.data[i], h.data[j]) < 0
}
