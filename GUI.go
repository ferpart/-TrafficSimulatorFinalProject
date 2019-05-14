package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const screenWidth = 550
const screenHeight = 550

var listCars map[string]Position
var carsInMap map[int]Position
var colorCar map[int]uint8
var s1 rand.Source

type Position struct {
	x, y, w, h int32
}

func GUI(cars []Car) {
	listCars = make(map[string]Position)
	carsInMap = make(map[int]Position)
	colorCar = make(map[int]uint8)

	s1 = rand.NewSource(time.Now().UnixNano())
	for i := 0; i < len(cars)*3; i = i + 3 {
		colorCar[i] = getRandColor()
		colorCar[i+1] = getRandColor()
		colorCar[i+2] = getRandColor()
	}

	// W - E
	listCars["60"] = Position{x: 20, y: 310, w: 20, h: 10}
	listCars["61"] = Position{x: 70, y: 310, w: 20, h: 10}
	listCars["62"] = Position{x: 120, y: 310, w: 20, h: 10}
	listCars["63"] = Position{x: 170, y: 310, w: 20, h: 10}
	listCars["64"] = Position{x: 220, y: 310, w: 20, h: 10}
	listCars["66"] = Position{x: 320, y: 310, w: 20, h: 10}
	listCars["67"] = Position{x: 370, y: 310, w: 20, h: 10}

	// S - N
	listCars["106"] = Position{x: 310, y: 510, w: 10, h: 20}
	listCars["96"] = Position{x: 310, y: 460, w: 10, h: 20}
	listCars["86"] = Position{x: 310, y: 410, w: 10, h: 20}
	listCars["76"] = Position{x: 310, y: 360, w: 10, h: 20}
	listCars["66"] = Position{x: 310, y: 310, w: 10, h: 20}
	listCars["46"] = Position{x: 310, y: 210, w: 10, h: 20}
	listCars["36"] = Position{x: 310, y: 160, w: 10, h: 20}

	// N - S
	listCars["04"] = Position{x: 230, y: 20, w: 10, h: 20}
	listCars["14"] = Position{x: 230, y: 70, w: 10, h: 20}
	listCars["24"] = Position{x: 230, y: 120, w: 10, h: 20}
	listCars["34"] = Position{x: 230, y: 170, w: 10, h: 20}
	listCars["44"] = Position{x: 230, y: 220, w: 10, h: 20}
	listCars["64"] = Position{x: 230, y: 320, w: 10, h: 20}
	listCars["74"] = Position{x: 230, y: 370, w: 10, h: 20}

	// E - W
	listCars["410"] = Position{x: 510, y: 230, w: 20, h: 10}
	listCars["49"] = Position{x: 460, y: 230, w: 20, h: 10}
	listCars["48"] = Position{x: 410, y: 230, w: 20, h: 10}
	listCars["47"] = Position{x: 360, y: 230, w: 20, h: 10}
	listCars["46"] = Position{x: 310, y: 230, w: 20, h: 10}
	listCars["44"] = Position{x: 210, y: 230, w: 20, h: 10}
	listCars["43"] = Position{x: 160, y: 230, w: 20, h: 10}

	for i, v := range cars {
		carsInMap[i] = listCars[v.originPos]
	}

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		os.Exit(1)
	}

	window, err := sdl.CreateWindow("City Traffic Simulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	renderer.Clear()

	var event sdl.Event
	isRunning := true

	// main loop
	for isRunning {
		// handle events, in this case escape key and close window
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					isRunning = false
				}
			}
		}

		// Creating the map
		// Set color gray
		// red, green, blue, alpha (alpha determines opaque-ness - usually 255)
		renderer.SetDrawColor(128, 128, 128, 255)

		renderer.Clear()

		// 1st quadrant
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&sdl.Rect{350, 0, 200, 200})

		// 2nd quadrant
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&sdl.Rect{0, 0, 200, 200})

		// 3rd quadrant
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&sdl.Rect{0, 350, 200, 200})

		// 4th quadrant
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&sdl.Rect{350, 350, 200, 200})
		renderer.SetDrawColor(0, 0, 255, 255)

		// Yellow Lines
		renderer.SetDrawColor(255, 255, 102, 255)
		renderer.DrawLine(275, 0, 275, 550)
		renderer.DrawLine(0, 275, 550, 275)

		// Semaphores
		// 1st quad
		if semaphoresInMap[1] == true {
			renderer.SetDrawColor(210, 255, 186, 255)
		} else {
			renderer.SetDrawColor(255, 0, 0, 255)
		}
		renderer.FillRect(&sdl.Rect{350, 200, 10, 35})

		// 2nd quad
		if semaphoresInMap[0] == true {
			renderer.SetDrawColor(210, 255, 186, 255)
		} else {
			renderer.SetDrawColor(255, 0, 0, 255)
		}
		renderer.FillRect(&sdl.Rect{200, 190, 35, 10})

		// 3rd quad
		if semaphoresInMap[3] == true {
			renderer.SetDrawColor(210, 255, 186, 255)
		} else {
			renderer.SetDrawColor(255, 0, 0, 255)
		}
		renderer.FillRect(&sdl.Rect{190, 315, 10, 35})

		// 4th quad
		if semaphoresInMap[2] == true {
			renderer.SetDrawColor(210, 255, 186, 255)
		} else {
			renderer.SetDrawColor(255, 0, 0, 255)
		}
		renderer.FillRect(&sdl.Rect{315, 350, 35, 10})

		for i, v := range carsInMap {
			renderer.SetDrawColor(colorCar[i], colorCar[i+1], colorCar[i+2], 255)
			renderer.FillRect(&sdl.Rect{v.x, v.y, v.w, v.h})
		}

		// The rects have been drawn, now it is time to tell the renderer to show
		// what has been draw to the screen
		renderer.Present()

		// quick and dirty way to do animation without taking into account how much
		// time has passed or that different computers will run at different speeds
		// and therefore the blue rect might be moving too fast on some computers to
		// see it, ruining the demonstration. Change the time value to experiment with
		// different blue rect speeds
		time.Sleep(time.Millisecond * 10)

	}

	renderer.Destroy()
	window.Destroy()

	sdl.Quit()
}

func getRandColor() uint8 {
	r1 := rand.New(s1)
	return uint8(r1.Intn(255))
}
