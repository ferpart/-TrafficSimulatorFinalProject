package main

import "flag"

func main() {

	ncars := flag.Int("ncars", 4, "cars to spawn")
	flag.Parse()

	city := City{}
	city.init()
	city.createMap(11)
	city.setSemaphores()
	city.run(*ncars)
}
