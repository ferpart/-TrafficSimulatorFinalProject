package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var pointsNorth map[Point]bool
var pointsEast map[Point]bool
var pointsSouth map[Point]bool
var pointsWest map[Point]bool
var pointToNode map[string]string

// City :
// Creates a city instance.
type City struct {
	cMap    [][]int
	mutex   *sync.Mutex
	semList []Semaphore
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

	pointsNorth = make(map[Point]bool, 0)
	pointsWest = make(map[Point]bool, 0)
	pointsSouth = make(map[Point]bool, 0)
	pointsEast = make(map[Point]bool, 0)

	pointToNode = make(map[string]string)

	pointsNorth[Point{0, 4}] = false
	pointsNorth[Point{1, 4}] = false
	pointsNorth[Point{2, 4}] = false

	pointsWest[Point{6, 0}] = false
	pointsWest[Point{6, 1}] = false
	pointsWest[Point{6, 2}] = false

	pointsSouth[Point{8, 6}] = false
	pointsSouth[Point{9, 6}] = false
	pointsSouth[Point{10, 6}] = false

	pointsEast[Point{4, 8}] = false
	pointsEast[Point{4, 9}] = false
	pointsEast[Point{4, 10}] = false

	pointToNode["04"] = "northC"
	pointToNode["14"] = "northB"
	pointToNode["24"] = "northA"

	pointToNode["60"] = "westC"
	pointToNode["61"] = "westB"
	pointToNode["62"] = "westA"

	pointToNode["86"] = "southA"
	pointToNode["96"] = "southB"
	pointToNode["106"] = "southC"

	pointToNode["48"] = "eastA"
	pointToNode["49"] = "eastB"
	pointToNode["410"] = "eastC"
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
	listPoint := make([]Point, 0)
	for i := 0; i < cars; i++ {
		p, err := getRandomPointStart()
		if err != nil {
			continue
		}
		listPoint = append(listPoint, p)
	}

	Graph()
	g := getItemGraph()
	l := make([]string, 0)
	for _, p := range listPoint {
		c := Car{originPos: p}
		pointString := strconv.Itoa(p.x) + strconv.Itoa(p.y)
		nodeName := pointToNode[pointString]
		indexNode := getIndex(nodeName)
		l = append(l, nodeName)
		c.index = nodeName
		g.nodes[indexNode].setCar(&c)
		go g.nodes[indexNode].getCar().move()
	}
	/*
		fmt.Println(l[0])
		fmt.Println("southF") */
	fmt.Println(getPath("westC", "westF"))
	time.Sleep(5 * time.Second)
}

// Get a random point from available positions
func getRandomPointStart() (Point, error) {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r1.Intn(4)
	fmt.Println(i)
	switch i {
	case 0:
		for k, v := range pointsNorth {
			if !v {
				pointsNorth[k] = true
				return k, nil
			}
		}
		return Point{}, errors.New("didn't find a point available")
	case 1:
		for k, v := range pointsWest {
			if !v {
				pointsWest[k] = true
				return k, nil
			}
		}
		return Point{}, errors.New("didn't find a point available")
	case 2:
		for k, v := range pointsSouth {
			if !v {
				pointsSouth[k] = true
				return k, nil
			}
		}
		return Point{}, errors.New("didn't find a point available")
	case 3:
		for k, v := range pointsEast {
			if !v {
				pointsEast[k] = true
				return k, nil
			}
		}
		return Point{}, errors.New("didn't find a point available")
	}
	return Point{}, errors.New("didn't find a point available")
}

// Initialize the semaphores and cars, to move around the city.
func (c *City) run(cars int) {
	// Semaphores begin
	go c.semList[0].acquireTurn()
	go c.semList[1].acquireTurn()
	go c.semList[2].acquireTurn()
	go c.semList[3].acquireTurn()

	// Generate cars inside the graph
	c.generateCars(cars)

	// use group waits instead
	time.Sleep(3600 * time.Second)
}
