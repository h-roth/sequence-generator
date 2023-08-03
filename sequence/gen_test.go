package sequence

import (
	"testing"
)

func TestGetNextSeqMutex(t *testing.T) {
	sequence = 0
	for i := int64(1); i <= 1000; i++ {
		if seq := getNextSeqMutex(); seq != i {
			t.Errorf("Expected %d, got %d", i, seq)
		}
	}
}

func TestGetNextSeqAtomic(t *testing.T) {
	sequence = 0
	for i := int64(1); i <= 1000; i++ {
		if seq := getNextSeqAtomic(); seq != i {
			t.Errorf("Expected %d, got %d", i, seq)
		}
	}
}
