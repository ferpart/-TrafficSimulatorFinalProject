package main

import (
	"time"
)

// Car :
// Creates a car instance.
type Car struct {
	originPos      Point // Origin point
	currentPos     Point // Current point
	destinationPos Point // Destination point
	velocity       float32
	index          string
	// cMap           *City
}

func (c *Car) move(path []string) {
	lastNode := g.nodes[getIndex(path[len(path)-1])]
	for {
		// is my velocity enough to move?
		if len(path) == 0 {
			g.lock.Lock()
			lastNode.setHasCar(false)
			g.lock.Unlock()
			return
		}
		if c.velocity >= 1 {
			g.lock.Lock()
			currentNode := g.nodes[getIndex(path[0])]
			nextNode := g.nodes[getIndex(path[1])]
			g.lock.Unlock()
			path = path[1:]
			// is next position a semaphor?
			if nextNode.getIsSemaphor() {
				// semaphor is in red -> semaphorState = false
				if !nextNode.getSemaphorState() {
					c.velocity = 0
					continue
				}
				// check if someone infront
				if nextNode.getHasCar() {
					c.velocity = 0
					continue
				}

				// move
				g.lock.Lock()
				currentNode.setHasCar(false)
				nextNode.setHasCar(true)
				g.lock.Unlock()
			} else {
				// is not semaphore, check if someone is infront
				if nextNode.getHasCar() {
					c.velocity = 0
					continue
				}
				g.lock.Lock()
				currentNode.setHasCar(false)
				nextNode.setHasCar(true)
				g.lock.Unlock()
			}
		}
		time.Sleep(1 * time.Second)
		c.velocity += c.velocity
	}
}
