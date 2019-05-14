package main

import (
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

func (s *Semaphore) acquireTurn(id int, name string) {
	semaphoreNode := g.nodes[getIndex(name)]
	for {
		s.mutex.Lock()
		semaphoresInMap[id] = true
		// fmt.Printf("\t\t==== SEMAPHORE %s ====\n", name)
		semaphoreNode.setSemaphorState(true)
		time.Sleep(2 * time.Second)
		s.mutex.Unlock()
		semaphoresInMap[id] = false
		semaphoreNode.setSemaphorState(false)
		time.Sleep(6 * time.Second)
	}

}
