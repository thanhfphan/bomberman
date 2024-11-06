package engine

type BitSet struct {
	bits []uint64
}

func NewBitSet() *BitSet {
	return &BitSet{
		bits: make([]uint64, 0),
	}
}

func (bs *BitSet) ensureCapacity(slot int) {
	if slot >= len(bs.bits) {
		newSize := len(bs.bits)*2 + 1 // +1 to avoid doubling to 0
		if newSize <= slot {
			newSize = slot + 1
		}

		newBits := make([]uint64, newSize)
		copy(newBits, bs.bits)
		bs.bits = newBits
	}
}

func (bs *BitSet) Add(value int) {
	if value < 0 {
		return
	}
	slot, pos := value/64, value%64
	bs.ensureCapacity(slot)
	bs.bits[slot] |= 1 << pos
}

func (bs *BitSet) Remove(value int) {
	if value < 0 {
		return
	}
	slot, pos := value/64, value%64
	if slot >= len(bs.bits) {
		return // Value is beyond the current capacity
	}
	bs.bits[slot] &^= 1 << pos
}

func (bs *BitSet) Contains(value int) bool {
	if value < 0 {
		return false
	}
	slot, pos := value/64, value%64
	if slot >= len(bs.bits) {
		return false // Value is beyond the current capacity
	}
	return (bs.bits[slot] & (1 << pos)) != 0
}
