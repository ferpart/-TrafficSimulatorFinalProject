package main

// Intersection :
// Creates a instersection instance.
type Intersection struct {
	listSemaphores []Semaphore
	turn           chan int
	// Should save who has the turn right now
}

func (i *Intersection) init() {
	// Create unique turn per semaphore
	i.turn = make(chan int, 1)
}

func (i *Intersection) offerTurn(n int) {

}
