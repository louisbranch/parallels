package main

import (
	"github.com/luizbranco/parallels/input"
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

const FPS = 60
const TileSize = 30

var window *sdl.Window
var renderer *sdl.Renderer
var mode Mode

type Vector3 struct {
	X int32
	Y int32
	Z int32
}

var Ship = &Vector3{50, 50, 0}

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

	image, err := img.Load("assets/images/ship.png")

	if err != nil {
		panic(err)
	}

	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(err)
	}

	defer texture.Destroy()

	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}

	if input.UpKey == input.KeyPressed || input.UpKey == input.KeyHeld {
		Ship.Y--
	}

	if input.DownKey == input.KeyPressed || input.DownKey == input.KeyHeld {
		Ship.Y++
	}

	if input.LeftKey == input.KeyPressed || input.LeftKey == input.KeyHeld {
		Ship.X--
	}

	if input.RightKey == input.KeyPressed || input.RightKey == input.KeyHeld {
		Ship.X++
	}

	src := sdl.Rect{W: w, H: h}
	dst := sdl.Rect{W: w, H: h, X: Ship.X, Y: Ship.Y}

	// Set renderer to black color (RGBA)
	renderer.SetDrawColor(0, 0, 0, 255)

	// Clear renderer to draw color
	renderer.Clear()

	rect := sdl.Rect{X: 0, Y: 0, W: 100, H: 100}

	for i := 0; i < 11; i++ {
		for j := 0; j < 30; j++ {
			if j%2 == 0 {
				renderer.SetDrawColor(0, 0, 255, 255)
			} else {
				renderer.SetDrawColor(255, 0, 0, 255)
			}
			renderer.FillRect(&rect)
			rect.X += 100
		}
		rect.X = 0
		rect.Y += 100
	}

	// Display Ship
	renderer.Copy(texture, &src, &dst)

	// Display render at the window
	renderer.Present()
}

func draw() {
	// Set renderer to black color (RGBA)
	renderer.SetDrawColor(0, 0, 0, 255)

	// Clear renderer to draw color
	renderer.Clear()

	// Display render at the window
	renderer.Present()
}
