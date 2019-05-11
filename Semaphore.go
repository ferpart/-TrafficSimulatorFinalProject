package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore :
// Creates a semaphore instance.
type Semaphore struct {
	// carQueue Queue
	id       int
	position Point
	mutex    *sync.Mutex
}

func (s *Semaphore) acquireTurn() {
	for {
		s.mutex.Lock()
		fmt.Printf("Semaphore %d has the turn\n", s.id)
		time.Sleep(2 * time.Second)
		s.mutex.Unlock()
		time.Sleep(6 * time.Second)
	}

}
