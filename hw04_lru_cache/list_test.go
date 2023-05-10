package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("complex2", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)       // [10]
		l.PushBack(20)        // [10, 20]
		l.PushFront(30)       // [30, 10, 20]
		elem := l.Back().Prev // 10
		l.Remove(elem)        // [30, 20]
		require.Equal(t, 2, l.Len())

		// require.Equal(t, (*ListItem)(nil), l.Back().Next)
		require.Nil(t, l.Back().Next)

		require.Equal(t, 30, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)

		l.MoveToFront(l.Back()) // [20, 30]

		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)

		for idx := 0; idx < 5; idx++ {
			l.PushFront(idx)
		} // [4, 3, 2, 1, 0, 20, 30]

		l.Remove(l.Front().Next)      // [4, 2, 1, 0, 20, 30]
		l.MoveToFront(l.Front().Next) // [2, 4, 1, 0, 20, 30]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}

		require.Equal(t, []int{2, 4, 1, 0, 20, 30}, elems)
	})
}
