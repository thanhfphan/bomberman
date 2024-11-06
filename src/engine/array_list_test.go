package engine

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
