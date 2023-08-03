package sequence

import (
	"log"
	"sync"
	"sync/atomic"
)

// TODO Probably avoid global vars
var (
	sequence int64
	mutex    sync.Mutex
)

func PublicGetNextID() int64 {
	nextID := getNextSeqMutex()
	// nextID := getNextSeqAtomic()
	log.Printf("Next ID: %d\n", nextID)
	return nextID
}

func getNextSeqAtomic() int64 {
	return atomic.AddInt64(&sequence, 1)
}

func getNextSeqMutex() int64 {
	mutex.Lock()
	defer mutex.Unlock()

	sequence++
	return sequence
}
