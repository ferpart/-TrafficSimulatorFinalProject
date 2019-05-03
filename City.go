package main

import "fmt"

// City :
// Creates a city instance.
type City struct {
	cMap        [][]int
	cSemaphores []Semaphore
}

// Creates a NxM matrix that will be the city map.
func (c *City) createMap(n, m int) {
	c.cMap = make([][]int, n, n)
	for i := range c.cMap {
		c.cMap[i] = make([]int, m, m)
	}
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
			if (c.cMap)[i][j] != 0 {
				fmt.Printf(" \033[1;31m%d\033[0m ", (c.cMap)[i][j])
			} else {
				fmt.Printf(" %d ", (c.cMap)[i][j])
			}
		}
	}
}

// Adds a new semaphore to the city.
func (c *City) addSemaphore(s Semaphore) {
	c.cSemaphores = append(c.cSemaphores, s)
}

// Prints city's semaphores in a human readable way.
func (c *City) printSemaphores() {
	fmt.Printf("\nCitySemaphores:\n")
	for _, v := range c.cSemaphores {
		fmt.Printf("\tSemaphore")
		fmt.Println(v)
	}
}
