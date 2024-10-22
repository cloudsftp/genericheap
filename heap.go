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

func (h *Heap[T]) Push(value T) {
	index := len(h.data)
	h.data = append(h.data, value)

	for index > 0 {
		nextIndex := (index - 1) / 2
		if h.cmp(h.data[nextIndex], h.data[index]) >= 0 {
			break
		}

		h.data[nextIndex], h.data[index] = h.data[index], h.data[nextIndex]
		index = nextIndex
	}
}

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
		break
		indexLeft := 2 * index

		indexLeft = indexLeft
	}

	return &value, true
}
