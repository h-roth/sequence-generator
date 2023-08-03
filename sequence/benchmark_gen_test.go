package sequence

import (
	"testing"
)

func BenchmarkGetNextSeqMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNextSeqMutex()
	}
}

func BenchmarkParallelGetNextSeqMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			getNextSeqMutex()
		}
	})
}

func BenchmarkGetNextSeqAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNextSeqAtomic()
	}
}

func BenchmarkParallelGetNextSeqAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			getNextSeqAtomic()
		}
	})
}
