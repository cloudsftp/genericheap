package genericheap

import "errors"

type ComparisonFunc[T any] func(a, b T) bool

type Heap[T any] struct {
	data    []T
	compare ComparisonFunc[T]
}

func NewHeap[T any](compare ComparisonFunc[T]) Heap[T] {
	return Heap[T]{
		data:    []T{},
		compare: compare,
	}
}

func (h *Heap[T]) Push(value T) {
	index := len(h.data)
	h.data = append(h.data, value)

	for index > 0 {
		nextIndex := index / 2

		if !h.compare(h.data[index], h.data[nextIndex]) {
			break
		}

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
		indexLeft := 2 * index
	}

	return &value, true
}
