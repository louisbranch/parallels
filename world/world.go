package world

import (
	"image/color"
	"math/rand"
	"time"
)

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

type World [WorldWith * WorldHeight]Terrain

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

const WorldTiles = 400
const WorldWith = 16 * WorldTiles
const WorldHeight = 10 * WorldTiles

var Earth, Arcadian, Underworld World

func init() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < WorldWith*WorldHeight; i++ {
		Earth[i] = Terrain(rand.Intn(8))
	}

}
