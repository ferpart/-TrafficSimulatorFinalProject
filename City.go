package main

import (
	"fmt"
	"sync"
	"time"
)

// City :
// Creates a city instance.
type City struct {
	cMap    [][]int
	mutex   *sync.Mutex
	semList []Semaphore
	carList []Car
}

// Initializes variables inside instance.
func (c *City) init() {
	c.mutex = &sync.Mutex{}
}

// Creates a NxN matrix that will be the city map.
func (c *City) createMap(n int) {
	c.cMap = make([][]int, n, n)
	for i := range c.cMap {
		c.cMap[i] = make([]int, n, n)
	}
	width := (n - 3) / 2
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			c.cMap[i][j] = 1
		}
	}
	for i := width + 3; i < width*2+3; i++ {
		for j := 0; j < width; j++ {
			c.cMap[i][j] = 1
		}
	}
	for i := 0; i < width; i++ {
		for j := width + 3; j < width*2+3; j++ {
			c.cMap[i][j] = 1
		}
	}
	for i := width + 3; i < width*2+3; i++ {
		for j := width + 3; j < width*2+3; j++ {
			c.cMap[i][j] = 1
		}
	}

}

// Set up n semaphores in the city
func (c *City) setSemaphores(n int) {
	c.semList = append(c.semList, Semaphore{id: 0, position: Point{4, 7}, mutex: c.mutex})
	c.semList = append(c.semList, Semaphore{id: 1, position: Point{3, 4}, mutex: c.mutex})
	c.semList = append(c.semList, Semaphore{id: 2, position: Point{6, 3}, mutex: c.mutex})
	c.semList = append(c.semList, Semaphore{id: 3, position: Point{7, 6}, mutex: c.mutex})
	c.cMap[4][7] = 2
	c.cMap[3][4] = 2
	c.cMap[6][3] = 2
	c.cMap[7][6] = 2
}

// Prints the city map in a human readable way.
func (c *City) printMap() {
	fmt.Printf("    ")
	for i := range c.cMap {
		fmt.Printf("%3d", i%10)
	}
	for i := range c.cMap {
		fmt.Printf("\n%4d ", i)
		for j := range (c.cMap)[i] {
			if (c.cMap)[i][j] == 1 {
				fmt.Printf(" \033[41m\033[1;31m%d\033[0m ", (c.cMap)[i][j])
			} else if (c.cMap)[i][j] == 2 {
				fmt.Printf(" \033[42m\033[1;32m%d\033[0m ", (c.cMap)[i][j])
			} else {
				fmt.Printf(" %d ", (c.cMap)[i][j])
			}
		}
	}
	fmt.Println()
}

// Generates the Cars
func (c *City) generateCars(cars int) {

}

// Initialize the semaphores and cars, to move around the city.
func (c *City) run(cars int) {
	c.generateCars(cars)
	go c.semList[0].acquireTurn()
	go c.semList[1].acquireTurn()
	go c.semList[2].acquireTurn()
	go c.semList[3].acquireTurn()

	// use group waits instead
	time.Sleep(3600 * time.Second)
}
