package camera

import "github.com/luizbranco/parallels/math"

const minZoom = 1
const maxZoom = 3

type Camera struct {
	Zoom int
	math.Rect
}

func (c *Camera) ZoomIn() {
	c.Zoom = math.Clamp(c.Zoom-1, minZoom, maxZoom)
}

func (c *Camera) ZoomOut() {
	c.Zoom = math.Clamp(c.Zoom+1, minZoom, maxZoom)
}

func (c *Camera) Clip(width, height, pixels int) (start, w, h int) {
	// camera tile-based size
	w = math.Clamp(math.DivCeil(c.W, pixels), 0, width)
	h = math.Clamp(math.DivCeil(c.H, pixels), 0, height)

	// camera tile-based position
	x := math.Clamp(math.DivFloor(c.X, pixels), 0, width-w)
	y := math.Clamp(math.DivFloor(c.Y, pixels), 0, height-h)

	start = x + (y * width)

	return
}
