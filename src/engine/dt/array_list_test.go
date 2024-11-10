package dt

import (
	"testing"
)

const (
	TenMillion = 10_000_000
)

func BenchmarkArrayList_Append(b *testing.B) {
	list := NewArrayList[int](1)
	for i := 0; i < TenMillion; i++ {
		list.Append(i)
	}
}

func BenchmarkArrayList_Get(b *testing.B) {
	// f, err := os.Create("mem.prof")
	// if err != nil {
	// 	b.Fatalf("could not create memory profile: %v", err)
	// }
	// defer f.Close()

	list := NewArrayList[int](7)
	for i := 0; i < TenMillion; i++ {
		list.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Get(i % TenMillion)
	}

	// b.StopTimer()
	// if err := pprof.WriteHeapProfile(f); err != nil {
	// 	b.Fatalf("could not write memory profile: %v", err)
	// }

	// go tool pprof -http=":8080" mem.prof
}

func BenchmarkArrayList_Remove(b *testing.B) {
	list := NewArrayList[int](7)
	for i := 0; i < TenMillion; i++ {
		list.Append(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Remove(i % TenMillion)
	}

}

func TestArrayList_Append(t *testing.T) {
	list := NewArrayList[int](1)
	for i := 0; i < 100; i++ {
		index := list.Append(i)
		if index != i {
			t.Errorf("expected index %d, got %d", i, index)
		}
	}
}

func TestArrayList_Get(t *testing.T) {
	list := NewArrayList[int](1)
	for i := 0; i < 100; i++ {
		list.Append(i)
	}
	for i := 0; i < 100; i++ {
		value, err := list.Get(i)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if value != i {
			t.Errorf("expected value %d, got %d", i, value)
		}
	}
}

func TestArrayList_Remove(t *testing.T) {
	list := NewArrayList[int](1)
	for i := 0; i < 100; i++ {
		list.Append(i)
	}
	for i := 0; i < 100; i++ {
		err := list.Remove(i)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		_, err = list.Get(i)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	}
}
func TestArrayList_AppendRemoveFreeList(t *testing.T) {
	list := NewArrayList[int](1)
	for i := 0; i < 10; i++ {
		list.Append(i)
	}

	for i := 0; i < 5; i++ {
		err := list.Remove(i)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	// Check if removed indices are in the free list
	for i := 0; i < 5; i++ {
		if !list.freeList.Contains(i) {
			t.Errorf("expected index %d to be in the free list", i)
		}
	}

	// Append new elements and check if they use the free list
	for i := 0; i < 5; i++ {
		index := list.Append(i + 10)
		expected := 5 - 1 - i // We use stack for the last in first out
		if index != expected {
			t.Errorf("expected index %d, got %d", expected, index)
		}
	}

	if list.freeList.Size() != 0 {
		t.Errorf("expected free list to be empty, got size %d", list.freeList.Size())
	}

	// the free list should empty now
	for i := 0; i < 5; i++ {
		if list.freeList.Contains(i) {
			t.Errorf("expected empty in the free list but got %d", i)
		}
	}
}
