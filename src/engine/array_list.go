package engine

import "errors"

var ErrIndexOutOfRange = errors.New("index out of range")

type ArrayList[T any] struct {
	capacity int
	size     int
	items    []T
	freeList []int
}

func NewArrayList[T any](capacity int) *ArrayList[T] {
	if capacity < 1 {
		capacity = 1
	}
	return &ArrayList[T]{
		capacity: capacity,
		size:     0,
		items:    make([]T, capacity),
		freeList: make([]int, 0),
	}
}

func (list *ArrayList[T]) Append(item T) (int, error) {
	var index int
	if len(list.freeList) > 0 {
		index = list.freeList[0]
		list.freeList = list.freeList[1:]
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
	if index < 0 || index >= list.size || list.isFree(index) {
		var zero T
		return zero, ErrIndexOutOfRange
	}

	return list.items[index], nil
}

func (list *ArrayList[T]) Remove(index int) error {
	if index < 0 || index >= list.size || list.isFree(index) {
		return ErrIndexOutOfRange
	}

	list.freeList = append(list.freeList, index)
	var zero T
	list.items[index] = zero

	return nil
}

func (list *ArrayList[T]) Size() int {
	return list.size
}

func (list *ArrayList[T]) isFree(index int) bool {
	for _, i := range list.freeList {
		if i == index {
			return true
		}
	}

	return false
}
