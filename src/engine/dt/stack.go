package dt

type StackInt struct {
	items    []int
	size     int
	capacity int
	lookup   *BitSet
}

func NewStack(capacity int) *StackInt {
	if capacity < 1 {
		capacity = 1
	}
	return &StackInt{
		items:    make([]int, capacity),
		size:     0,
		capacity: capacity,
		lookup:   NewBitSet(),
	}
}

func (s *StackInt) Push(item int) {
	if s.size >= s.capacity {
		s.capacity *= 2
		newItems := make([]int, s.capacity)
		copy(newItems, s.items)
		s.items = newItems
	}
	s.items[s.size] = item
	s.lookup.Add(item)
	s.size++
}

func (s *StackInt) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	item := s.items[s.size]
	s.items[s.size] = 0
	s.size--
	s.lookup.Remove(item)
	return item, true
}

func (s *StackInt) Peek() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	return s.items[s.size], true
}

func (s *StackInt) Size() int {
	return s.size
}

func (s *StackInt) IsEmpty() bool {
	return s.size == 0
}

func (s *StackInt) Contains(item int) bool {
	return s.lookup.Contains(item)
}
