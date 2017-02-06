package world

import (
	"image/color"
	"math/rand"
	"time"
)

const TileSize = 50

type Terrain int

var TerrainColor = [...]color.RGBA{
	{0, 0, 255, 255},
	{0, 0, 100, 255},
	{0, 255, 0, 255},
	{200, 100, 100, 255},
	{100, 100, 100, 255},
	{0, 100, 0, 255},
	{150, 150, 150, 255},
	{255, 255, 255, 255},
}

type World struct {
	W       int
	H       int
	Terrain []Terrain
}

const (
	Water Terrain = iota
	DeepWater
	Land
	Grass
	Mountain
	Swamp
	Forest
	Desert
	Tundra
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (w *World) Build() {
	w.Terrain = make([]Terrain, w.Size())

	for i := range w.Terrain {
		w.Terrain[i] = Terrain(rand.Intn(8))
	}
}

func (w *World) Size() int {
	return w.W * w.H
}
