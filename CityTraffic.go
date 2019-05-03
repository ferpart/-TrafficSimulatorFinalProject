package main

func main() {
	city := City{}
	city.createMap(20, 20)
	city.printMap()

}

/*
sem1 := Semaforo{carQueue: Queue{queue: make([]Car, 0), len: 0}}
	sem1.carQueue.add(Car{x: 1, y: 2})
	sem1.carQueue.add(Car{x: 2, y: -21})
	sem1.carQueue.add(Car{x: 4, y: -31})
	sem1.carQueue.print()

	for {
		if _, err := sem1.carQueue.get(); err != nil {
			fmt.Println(err)
			break
		}
		sem1.carQueue.print()
	}
*/
