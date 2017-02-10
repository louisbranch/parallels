package main

import (
	"github.com/luizbranco/parallels/camera"
	"github.com/luizbranco/parallels/input"
	"github.com/luizbranco/parallels/math"
	"github.com/luizbranco/parallels/world"
	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

const FPS = 60

var window *sdl.Window
var renderer *sdl.Renderer
var cam *camera.Camera
var earth *world.World
var mode Mode

type Mode int

type Image struct {
	W       int32
	H       int32
	Texture *sdl.Texture
}

var water Image
var forest Image

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
	flags = sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC
	renderer, err = sdl.CreateRenderer(window, -1, flags)

	// If a rendered can't be created, we crash
	if err != nil {
		panic(err)
	}

	// Initialize global vals
	mode = GameMode
	cam = &camera.Camera{
		TileSize: 50,
		Speed:    10,
		MinZoom:  1,
		MaxZoom:  3,
		Zoom:     1,
	}
	earth = &world.World{
		W: 200,
		H: 150,
	}
	earth.Build()

	water = loadImage("water_a1")
	forest = loadImage("forest")

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

	water.Texture.Destroy()
	forest.Texture.Destroy()

	// Destroy global renderer
	renderer.Destroy()

	// Destroy global window
	window.Destroy()

	// Close subsystems and exit
	sdl.Quit()
}

func drawMenu() {
	if input.EscKey == input.KeyPressed {
		mode = GameMode
	}

	draw()
}

func drawGame() {
	width, height := window.GetSize()
	cam.W = width
	cam.H = height
	cam.MaxW = earth.W*cam.TileSize/cam.Zoom - cam.W/2
	cam.MaxH = earth.H*cam.TileSize/cam.Zoom - cam.H/2

	if input.EscKey == input.KeyPressed {
		mode = MenuMode
	}

	if input.UpKey == input.KeyPressed || input.UpKey == input.KeyHeld {
		cam.MoveUp()
	}

	if input.DownKey == input.KeyPressed || input.DownKey == input.KeyHeld {
		cam.MoveDown()
	}

	if input.LeftKey == input.KeyPressed || input.LeftKey == input.KeyHeld {
		cam.MoveLeft()
	}

	if input.RightKey == input.KeyPressed || input.RightKey == input.KeyHeld {
		cam.MoveRight()
	}

	if input.ZoomInKey == input.KeyPressed {
		cam.ZoomIn()
	}

	if input.ZoomOutKey == input.KeyPressed {
		cam.ZoomOut()
	}

	// Set renderer to black color (RGBA)
	renderer.SetDrawColor(0, 0, 0, 255)

	// Clear renderer to draw color
	renderer.Clear()

	start, w, h := cam.Clip(earth.W, earth.H)

	size := int32(math.DivCeil(cam.TileSize, cam.Zoom))

	rect := &sdl.Rect{W: size, H: size}
	src := &sdl.Rect{W: water.W, H: water.H}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			t := earth.Tiles[start+x]
			if t.Terrain == world.Water {
				renderer.Copy(water.Texture, src, rect)
			} else {
				renderer.Copy(forest.Texture, src, rect)
			}
			rect.X += size
		}
		rect.X = 0
		rect.Y += size
		start += earth.W
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

func loadImage(name string) Image {
	image, err := img.Load("assets/images/" + name + ".png")
	if err != nil {
		panic(err)
	}
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(err)
	}

	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}

	return Image{
		W:       w,
		H:       h,
		Texture: texture,
	}
}
