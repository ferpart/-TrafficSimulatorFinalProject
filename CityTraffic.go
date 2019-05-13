package main

func main() {

	city := City{}
	city.init()
	city.createMap(11)
	city.setSemaphores()
	city.run(1)
}
