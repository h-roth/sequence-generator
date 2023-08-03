package sequence

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type IDGenerator interface {
	NextID() int64
}

type Sequence struct {
	generator IDGenerator
}

func New(generator IDGenerator) *Sequence {
	return &Sequence{generator: generator}
}

func (s *Sequence) GetNextID() int64 {
	return s.generator.NextID()
}

type AtomicGenerator struct {
	sequence int64
}

func (a *AtomicGenerator) NextID() int64 {
	return atomic.AddInt64(&a.sequence, 1)
}

type MutexGenerator struct {
	sequence int64
	mutex    sync.Mutex
}

func (m *MutexGenerator) NextID() int64 {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.sequence++
	return m.sequence
}

type TimestampGenerator struct {
	lastTimestamp int64
	sequence      int64
	nodeID        int64
	mutex         sync.Mutex
}

func NewTimestampGenerator(nodeID int64) *TimestampGenerator {
	return &TimestampGenerator{nodeID: nodeID}
}

// Uses nano seconds as the timestamp.
// First 51 bits are the timestamp, next 7 bits are the node ID, and last 6 bits are the sequence number
// If you have more than 128 nodes, or if you need more than 64 unique IDs per nanosecond will need to change.
func (s *TimestampGenerator) NextID() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	timestamp := time.Now().UnixNano()

	// If current timestamp is smaller than last timestamp, the system clock has been turned back
	if timestamp < s.lastTimestamp {
		log.Fatalf("Invalid system clock")
	}

	// If current timestamp is the same as the last one, increase the sequence
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 0x3F // Max sequence number is 63

		// Sleep if sequence has reached the max number in the same nanosecond timestamp
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	return (timestamp << 13) | (s.nodeID << 6) | s.sequence
}
