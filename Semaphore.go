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

func (s *Semaphore) acquireTurn(name string) {
	semaphoreNode := g.nodes[getIndex(name)]
	for {
		s.mutex.Lock()
		fmt.Printf("\t\t==== SEMAPHORE %s ====\n", name)
		semaphoreNode.setSemaphorState(true)
		time.Sleep(2 * time.Second)
		s.mutex.Unlock()
		semaphoreNode.setSemaphorState(false)
		time.Sleep(6 * time.Second)
	}

}
