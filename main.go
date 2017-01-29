package main

import (
	"github.com/luizbranco/parallels/input"
	"github.com/veandco/go-sdl2/sdl"
)

const FPS = 60

var window *sdl.Window
var renderer *sdl.Renderer
var mode Mode

type Mode int

const (
	MenuMode Mode = iota
	GameMode
)

func main() {
	// Initialize Video, Keyboard and Mouse Subsystems
	sdl.Init(sdl.INIT_VIDEO)

	// Set hardware accelaration, render is synchronized with refresh rate of the
	// screen, and window to fullscreen
	flags := uint32(sdl.RENDERER_ACCELERATED |
		sdl.RENDERER_PRESENTVSYNC |
		sdl.WINDOW_FULLSCREEN_DESKTOP)

	var err error

	// Create a fullscreen window without a definied position or size
	window, err = sdl.CreateWindow("Parallels", 0, 0, 0, 0, flags)

	// If a window can't be created, we crash
	if err != nil {
		panic(err)
	}

	// Create a renderer for the window. -1 is set to give us the first video
	// driver available. No flags (0) for the render is passed.
	renderer, err = sdl.CreateRenderer(window, -1, 0)

	// If a rendered can't be created, we crash
	if err != nil {
		panic(err)
	}

	mode = MenuMode

	// delay between loop interactions
	delay := uint32(1000 / FPS)

	for {
		start := sdl.GetTicks()

		input.Process()

		if input.QuitKey == input.KeyPressed {
			break
		}

		switch mode {
		case MenuMode:
			drawMenu()
		case GameMode:
			drawGame()
		}

		input.Update()

		last := sdl.GetTicks() - start

		if last < delay {
			sdl.Delay(delay - last)
		}
	}

	// Destroy global renderer
	renderer.Destroy()

	// Destroy global window
	window.Destroy()

	// Close subsystems and exit
	sdl.Quit()
}

func drawMenu() {
	if input.NextTurnKey == input.KeyPressed {
		mode = GameMode
	}

	draw()
}

func drawGame() {
	if input.NextTurnKey == input.KeyPressed {
		mode = MenuMode
	}

	// Set renderer to white color (RGBA)
	renderer.SetDrawColor(255, 255, 255, 255)

	// Clear renderer to draw color
	renderer.Clear()

	// Display render at the window
	renderer.Present()
}

func draw() {
	// Set renderer to white color (RGBA)
	renderer.SetDrawColor(0, 0, 0, 255)

	// Clear renderer to draw color
	renderer.Clear()

	// Display render at the window
	renderer.Present()
}
