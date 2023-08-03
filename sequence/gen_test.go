package sequence

import (
	"sync"
	"testing"
)

func TestSequence_GetNextID(t *testing.T) {
	g := &MutexGenerator{}
	s := New(g)

	for i := int64(1); i <= 10; i++ {
		if id := s.GetNextID(); id != i {
			t.Errorf("Expected %d, got %d", i, id)
		}
	}
}

func TestAtomicGenerator_NextID(t *testing.T) {
	g := &AtomicGenerator{}
	for i := int64(1); i <= 10; i++ {
		if id := g.NextID(); id != i {
			t.Errorf("Expected %d, got %d", i, id)
		}
	}
}

func TestMutexGenerator_NextID(t *testing.T) {
	g := &MutexGenerator{}
	for i := int64(1); i <= 10; i++ {
		if id := g.NextID(); id != i {
			t.Errorf("Expected %d, got %d", i, id)
		}
	}
}

func TestTimestampGenerator_NextID(t *testing.T) {
	g := NewTimestampGenerator(1)
	var lastID int64
	var m sync.Mutex
	var wg sync.WaitGroup

	// Simulate concurrent requests
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			id := g.NextID()

			m.Lock()
			if id <= lastID {
				t.Errorf("NextID should be greater than the last one")
			}
			lastID = id
			m.Unlock()
		}()
	}

	wg.Wait()
}
