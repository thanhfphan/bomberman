package engine

type BitSet struct {
	bits []uint64
}

func NewBitSet(maxValue int) *BitSet {
	numSlots := (maxValue + 63) / 64 // Each uint64 holds 64 bits
	return &BitSet{
		bits: make([]uint64, numSlots),
	}
}

func (bs *BitSet) Add(value int) {
	if value < 0 {
		return
	}
	slot, pos := value/64, value%64
	bs.bits[slot] |= 1 << pos
}

func (bs *BitSet) Remove(value int) {
	if value < 0 {
		return
	}
	slot, pos := value/64, value%64
	bs.bits[slot] &^= 1 << pos
}

func (bs *BitSet) Contains(value int) bool {
	if value < 0 {
		return false
	}
	slot, pos := value/64, value%64
	return (bs.bits[slot] & (1 << pos)) != 0
}
