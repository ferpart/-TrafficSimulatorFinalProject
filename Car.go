package main

import (
	"fmt"
	"time"
)

// Car :
// Creates a car instance.
type Car struct {
	id        int
	originPos string // Origin point
	velocity  float32
	index     string
	// cMap           *City
}

func (c *Car) move(id int, path []string, vel float32) {
	lastNode := g.nodes[getIndex(path[len(path)-1])]
	copyPath := path
	for {
		if len(path) <= 1 {
			g.lock.Lock()
			lastNode.setHasCar(false)
			carsInMap[id] = Position{600, 600, 1, 1}
			g.lock.Unlock()
			fmt.Printf("|%d| finished\n", id)
			fmt.Printf("|%d| %v\n", id, copyPath)
			return
		}
		// is my velocity enough to move?
		fmt.Printf("|%d| velocity: %f\n", id, c.velocity)
		if c.velocity >= 1 {
			g.lock.Lock()
			currentNode := g.nodes[getIndex(path[0])]
			nextNode := g.nodes[getIndex(path[1])]
			g.lock.Unlock()
			// is next position a semaphor?
			if currentNode.getIsSemaphor() {
				// semaphor is in red -> semaphorState = false
				if !currentNode.getSemaphorState() {
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
				// fmt.Printf("|%d| [%s] --> [%s] \t [%v]\n", id, path[0], path[1], path)
				currentNode.setHasCar(false)
				nextNode.setHasCar(true)
				carsInMap[id] = listCars[getMatrixPos[path[1]]]
				path = path[1:]
				g.lock.Unlock()
				c.velocity = 0
			} else {
				// is not semaphore, check if someone is infront
				if nextNode.getHasCar() {
					c.velocity = 0
					continue
				}
				// move
				g.lock.Lock()
				// fmt.Printf("|%d| [%s] --> [%s] \t [%v]\n", id, path[0], path[1], path)
				currentNode.setHasCar(false)
				nextNode.setHasCar(true)
				carsInMap[id] = listCars[getMatrixPos[path[1]]]
				path = path[1:]
				g.lock.Unlock()
				c.velocity = 0
			}
		}
		time.Sleep(500 * time.Millisecond)
		c.velocity += vel
	}
}
