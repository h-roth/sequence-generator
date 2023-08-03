package sequence

import "testing"

func BenchmarkMutexGenerator_NextID(b *testing.B) {
	g := &MutexGenerator{}
	for i := 0; i < b.N; i++ {
		g.NextID()
	}
}

func BenchmarkParallelMutexGenerator_NextID(b *testing.B) {
	g := &MutexGenerator{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.NextID()
		}
	})
}

func BenchmarkAtomicGenerator_NextID(b *testing.B) {
	g := &AtomicGenerator{}
	for i := 0; i < b.N; i++ {
		g.NextID()
	}
}

func BenchmarkParallelAtomicGenerator_NextID(b *testing.B) {
	g := &AtomicGenerator{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.NextID()
		}
	})
}

func BenchmarkTimestampGenerator_NextID(b *testing.B) {
	g := NewTimestampGenerator(1)
	for i := 0; i < b.N; i++ {
		g.NextID()
	}
}

func BenchmarkParallelTimestampGenerator_NextID(b *testing.B) {
	g := NewTimestampGenerator(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			g.NextID()
		}
	})
}
