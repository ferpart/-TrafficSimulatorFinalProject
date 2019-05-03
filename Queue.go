package main

import (
	"errors"
	"fmt"
)

// Queue :
// Simple implementation of a queue using slices.
type Queue struct {
	// This is expected to see the cars in the waiting line
	// and in the future using AI, decide if it's urgent to take the next turn
	// maybe a counter will do the same approach, we'll see
	queue []Car // TODO: Check if necessary
	Len   int   // Lenght of queue
}

// Adds a Car at the end of the queue.
func (q *Queue) add(car Car) {
	q.queue = append(q.queue, car)
	q.Len++
}

// Gets the first element in the queue.
func (q *Queue) get() (car Car, err error) {
	if q.Len < 1 {
		return Car{}, errors.New("queue is empty")
	}
	car = q.queue[0]
	q.queue = q.queue[1:]
	q.Len = q.Len - 1
	return car, nil
}

// Prints the elements in the Queue.
func (q *Queue) print() {
	fmt.Printf("\nQueue:\n")
	for i, v := range q.queue {
		fmt.Printf("\t[%d]: Car", i)
		fmt.Println(v)
	}
}
