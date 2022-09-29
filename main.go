package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	screenWidth          = 1000
	screenHeight         = 700
	targetTicksPerSecond = 60
)

type vector struct {
	x float64
	y float64
}

var delta float64

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	// Create Window
	window, err := sdl.CreateWindow("Shoten Game by Manraj", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer func(window *sdl.Window) {
		if err = window.Destroy(); err != nil {
			return
		}
	}(window)

	// Create Renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer func(renderer *sdl.Renderer) {
		if err = renderer.Destroy(); err != nil {
			return
		}
	}(renderer)

	elements = append(elements, newPlayer(renderer)) // Player Defined
	for i := 0; i < 8; i++ {                         // Enemy Define
		for j := 0; j < 3; j++ {
			x := (float64(i)/8)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)
			elements = append(elements, newBasicEnemy(renderer, x, y))
		}
	}
	initBulletPool(renderer)

	for {
		frameStartTime := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		err := renderer.SetDrawColor(255, 255, 255, 255)
		if err != nil {
			return
		}
		err2 := renderer.Clear()
		if err2 != nil {
			return
		}
		for _, elem := range elements {
			if elem.active {
				err = elem.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
				err = elem.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", err)
				}
			}
		}
		if err := checkCollisions(); err != nil {
			fmt.Println("checking collisions:", err)
			return
		}
		renderer.Present()
		delta = time.Since(frameStartTime).Seconds() * targetTicksPerSecond
	}
}
