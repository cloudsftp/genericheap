package genericheap

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntHeap(t *testing.T) {
	values := []int{1, 20, -1, 10}

	cmp := func(a, b int) int {
		return a - b
	}

	heap := NewHeap(cmp)

	for _, value := range values {
		heap.Push(value)
	}

	slices.SortFunc(values, cmp)

	for _, expected := range values {
		actual, ok := heap.Pop()
		require.True(t, ok)
		require.Equal(t, expected, *actual)
	}
}
