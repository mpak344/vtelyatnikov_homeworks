package doublylinkedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TetsLen(t *testing.T) {
	list := new(List)
	list.PushFront(123)
	list.PushFront(234)
	list.PushFront(345)
	require.Equal(t, list.Len(), 3)

	list2 := new(List)
	list2.PushBack(123)
	list2.PushBack(234)
	list2.PushBack(345)
	require.Equal(t, list2.Len(), 3)

	require.Equal(t, new(List).Len(), 0)
}

func TestPushFront(t *testing.T) {
	list := new(List)
	list.PushFront(123)
	require.Equal(t, list.First().Value(), 123)
	require.Equal(t, list.Last().Value(), 123)
	list.PushFront(234)
	require.Equal(t, list.First().Value(), 234)
	require.Equal(t, list.Last().Value(), 123)

	bigList := new(List)
	for i := 0; i < 10; i++ {
		bigList.PushFront(i)
	}
	require.Equal(t, bigList.First().Value(), 9)
	require.Equal(t, bigList.Last().Value(), 0)
}

func TestPushBack(t *testing.T) {
	list := new(List)
	list.PushBack(123)
	require.Equal(t, list.First().Value(), 123)
	require.Equal(t, list.Last().Value(), 123)
	list.PushBack(234)
	require.Equal(t, list.First().Value(), 123)
	require.Equal(t, list.Last().Value(), 234)

	bigList := new(List)
	for i := 0; i < 10; i++ {
		bigList.PushBack(i)
	}
	require.Equal(t, bigList.First().Value(), 0)
	require.Equal(t, bigList.Last().Value(), 9)
}

func TestRemove(t *testing.T) {
	list := new(List)
	for i := 0; i < 10; i++ {
		list.PushBack(i)
	}
	require.Equal(t, list.First().Value(), 0)
	list.Remove(*list.First())
	require.Equal(t, list.First().Value(), 1)
	require.Equal(t, list.Len(), 9)

	list.Remove(*list.Last())
	require.Equal(t, list.Last().Value(), 8)
	require.Equal(t, list.Len(), 8)

	for list.First() != nil {
		list.Remove(*list.First())
	}
	require.Equal(t, list.Len(), 0)
}
