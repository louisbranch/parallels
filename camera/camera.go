package camera

import "github.com/luizbranco/parallels/math"
import m "math"

const minZoom = 1
const maxZoom = 3

type Camera struct {
	math.Rect
	Zoom int
}

func (c *Camera) ZoomIn() {
	c.Zoom = math.Clamp(c.Zoom-1, minZoom, maxZoom)
}

func (c *Camera) ZoomOut() {
	c.Zoom = math.Clamp(c.Zoom+1, minZoom, maxZoom)
}

func (c *Camera) Clip(src math.Rect, size float64) (dst math.Rect) {
	dst.X = int(m.Ceil(float64(src.X) / size))
	dst.Y = int(m.Ceil(float64(src.Y) / size))
	dst.W = int(m.Ceil(float64(src.W) / size))
	dst.H = int(m.Ceil(float64(src.H) / size))

	return dst
}
