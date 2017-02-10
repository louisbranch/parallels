package world

import (
	"math/rand"
	"time"

	"github.com/luizbranco/parallels/world/perlin"
)

type World struct {
	W     int
	H     int
	Tiles []Tile
}

type Tile struct {
	Terrain
}

type Terrain int

const (
	Water Terrain = iota
	DeepWater
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
	p := perlin.New2D(0)

	w.Tiles = make([]Tile, w.W*w.H)

	width := float64(w.W)
	height := float64(w.H)

	for y := 0; y < w.H; y++ {
		for x := 0; x < w.W; x++ {

			maxW := width / 2.0
			maxH := height / 2.0

			nx := maxW - float64(x)/maxW
			ny := maxH - float64(y)/maxH

			noise := p.Get2D

			// Adapted from http://www.redblobgames.com/maps/terrain-from-noise/
			elevation := noise(1*nx, 1*ny) +
				0.50*noise(2*nx, 2*ny) +
				0.25*noise(4*nx, 4*ny) +
				0.13*noise(8*nx, 8*ny) +
				0.06*noise(16*nx, 16*ny) +
				0.03*noise(32*nx, 32*ny)

			moisture := 1.00*noise(1*nx, 1*ny) +
				0.75*noise(2*nx, 2*ny) +
				0.33*noise(4*nx, 4*ny) +
				0.33*noise(8*nx, 8*ny) +
				0.33*noise(16*nx, 16*ny) +
				0.50*noise(32*nx, 32*ny)

			magic := 1.00*noise(1*nx, 1*ny) +
				0.65*noise(2*nx, 2*ny) +
				0.40*noise(4*nx, 4*ny) +
				0.25*noise(8*nx, 8*ny) +
				0.15*noise(16*nx, 16*ny) +
				0.10*noise(32*nx, 32*ny)

			terrain := biome(elevation, moisture, magic)

			w.Tiles[x+y*w.W] = Tile{Terrain: terrain}
		}
	}
}

func biome(elevation, moisture, magic float64) Terrain {
	if elevation < 0 {
		return Water
	}

	return Grass
}
