package camera

import "github.com/luizbranco/parallels/math"

const MinZoom = 1
const MaxZoom = 3

type Camera struct {
	Zoom     int
	MinZoom  int
	MaxZoom  int
	Speed    int
	TileSize int
	MaxW     int
	MaxH     int
	math.Rect
}

func (c *Camera) MoveUp() {
	c.Y = math.Clamp(c.Y-c.Speed, 0, c.MaxH)
}

func (c *Camera) MoveDown() {
	c.Y = math.Clamp(c.Y+c.Speed, 0, c.MaxH)
}

func (c *Camera) MoveLeft() {
	c.X = math.Clamp(c.X-c.Speed, 0, c.MaxW)
}

func (c *Camera) MoveRight() {
	c.X = math.Clamp(c.X+c.Speed, 0, c.MaxW)
}

func (c *Camera) ZoomIn() {
	c.Zoom = math.Clamp(c.Zoom-1, c.MinZoom, c.MaxZoom)
}

func (c *Camera) ZoomOut() {
	c.Zoom = math.Clamp(c.Zoom+1, c.MinZoom, c.MaxZoom)
}

func (c *Camera) Clip(width, height int) (start, w, h int) {
	if c.Zoom == 0 {
		c.Zoom = 1
	}

	size := math.DivCeil(c.TileSize, c.Zoom)

	// camera tile-based size
	w = math.Clamp(math.DivCeil(c.W, size), 0, width)
	h = math.Clamp(math.DivCeil(c.H, size), 0, height)

	// camera tile-based position
	x := math.Clamp(math.DivFloor(c.X, size), 0, width-w)
	y := math.Clamp(math.DivFloor(c.Y, size), 0, height-h)

	start = x + (y * width)

	return
}
