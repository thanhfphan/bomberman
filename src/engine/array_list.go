package engine

import "errors"

var ErrIndexOutOfRange = errors.New("index out of range")

type ArrayList[T any] struct {
	capacity int
	size     int
	items    []T
	freeList *StackInt
}

func NewArrayList[T any](capacity int) *ArrayList[T] {
	if capacity < 1 {
		capacity = 1
	}
	return &ArrayList[T]{
		capacity: capacity,
		size:     0,
		items:    make([]T, capacity),
		freeList: NewStack(capacity),
	}
}

func (list *ArrayList[T]) Append(item T) (int, error) {
	var index int
	if list.freeList.Size() > 0 {
		index, _ = list.freeList.Pop()
	} else {
		if list.size >= list.capacity {
			list.capacity *= 2
			newItems := make([]T, list.capacity)
			copy(newItems, list.items)
			list.items = newItems
		}

		index = list.size
		list.size++
	}

	list.items[index] = item
	return index, nil
}

func (list *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= list.size || list.freeList.Contains(index) {
		return *new(T), ErrIndexOutOfRange
	}

	return list.items[index], nil
}

func (list *ArrayList[T]) Remove(index int) error {
	if index < 0 || index >= list.size || list.freeList.Contains(index) {
		return ErrIndexOutOfRange
	}

	list.freeList.Push(index)
	list.items[index] = *new(T)

	return nil
}

func (list *ArrayList[T]) Size() int {
	return list.size
}
