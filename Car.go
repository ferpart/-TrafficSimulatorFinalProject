package main

// Car :
// Creates a car instance.
type Car struct {
	originPos      Point // Origin point
	currentPos     Point // Current point
	destinationPos Point // Destination point
	velocity       float32
	cMap           *City
}
