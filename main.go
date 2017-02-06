package main

import (
	"github.com/luizbranco/parallels/input"
	"github.com/luizbranco/parallels/world"
	"github.com/veandco/go-sdl2/sdl"
)

const FPS = 60

var window *sdl.Window
var renderer *sdl.Renderer
var camera = &sdl.Rect{}
var earth = &world.World{}
var mode Mode

type Vector3 struct {
	X int32
	Y int32
	Z int32
}

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
	earth.Build()

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
	width, height := window.GetSize()
	camera.W = int32(width)
	camera.H = int32(height)

	const speed = world.TileSize

	if input.NextTurnKey == input.KeyPressed {
		mode = MenuMode
	}

	if input.UpKey == input.KeyPressed || input.UpKey == input.KeyHeld {
		camera.Y -= speed
		if camera.Y < 0 {
			camera.Y = 0
		}
	}

	if input.DownKey == input.KeyPressed || input.DownKey == input.KeyHeld {
		camera.Y += speed
		max := world.WorldHeight*world.TileSize - camera.H/2
		if camera.Y > max {
			camera.Y = max
		}
	}

	if input.LeftKey == input.KeyPressed || input.LeftKey == input.KeyHeld {
		camera.X -= speed
		if camera.X < 0 {
			camera.X = 0
		}
	}

	if input.RightKey == input.KeyPressed || input.RightKey == input.KeyHeld {
		camera.X += speed
		max := world.WorldWith*world.TileSize - camera.W/2
		if camera.X > max {
			camera.X = max
		}
	}

	// Set renderer to black color (RGBA)
	renderer.SetDrawColor(0, 0, 0, 255)

	// Clear renderer to draw color
	renderer.Clear()

	rect := sdl.Rect{X: 0, Y: 0, W: world.TileSize, H: world.TileSize}

	clip := earth.Intersect(*camera)

	for y := clip.Y * world.WorldWith; y < world.WorldLength; y += world.WorldWith {
		for x := clip.X; x < clip.X+clip.W; x++ {
			t := earth[x+y]
			color := world.TerrainColor[t]
			renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			renderer.FillRect(&rect)
			rect.X += world.TileSize
		}
		rect.X = 0
		rect.Y += world.TileSize
	}

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
