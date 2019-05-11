package main

func main() {

	city := City{}
	city.init()
	city.createMap(11)
	city.setSemaphores(4)
	city.printMap()
	city.generateCars(4)
	// city.run()

}
