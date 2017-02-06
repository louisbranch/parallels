package world

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
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

type World [WorldWith * WorldHeight]Terrain

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

const WorldTiles = 400
const WorldWith = 16 * WorldTiles
const WorldHeight = 10 * WorldTiles
const WorldLength = WorldWith * WorldHeight

func init() {
	rand.Seed(time.Now().UnixNano())

}

func (w *World) Build() {
	for i := 0; i < WorldLength; i++ {
		w[i] = Terrain(rand.Intn(8))
	}
}

func (w World) Intersect(src sdl.Rect) (dst sdl.Rect) {
	dst.X = int32(math.Ceil(float64(src.X) / TileSize))
	dst.Y = int32(math.Ceil(float64(src.Y) / TileSize))
	dst.W = int32(math.Ceil(float64(src.W) / TileSize))
	dst.H = int32(math.Ceil(float64(src.H) / TileSize))

	return dst
}
