package perlin

// Adapted from: https://gist.github.com/nsf/1170424

import (
	"math"
	"math/rand"
)

type vector2 struct {
	x, y float64
}

func lerp(a, b, v float64) float64 {
	return a*(1-v) + b*v
}

func smooth(v float64) float64 {
	return v * v * (3 - 2*v)
}

func randGradient(r *rand.Rand) vector2 {
	v := r.Float64() * math.Pi * 2
	return vector2{math.Cos(v), math.Sin(v)}
}

func gradient(orig, grad, p vector2) float64 {
	sp := vector2{p.x - orig.x, p.y - orig.y}
	return grad.x*sp.x + grad.y*sp.y
}

type Noise2D struct {
	rgradients   []vector2
	permutations []int
	gradients    [4]vector2
	origins      [4]vector2
}

func New2D(seed int) *Noise2D {
	rnd := rand.New(rand.NewSource(int64(seed)))

	noise := new(Noise2D)
	noise.rgradients = make([]vector2, 256)
	noise.permutations = rand.Perm(256)
	for i := range noise.rgradients {
		noise.rgradients[i] = randGradient(rnd)
	}

	return noise
}

func (noise *Noise2D) getGradient(x, y int) vector2 {
	idx := noise.permutations[x&255] + noise.permutations[y&255]
	return noise.rgradients[idx&255]
}

func (noise *Noise2D) getGradients(x, y float64) {
	x0f := math.Floor(x)
	y0f := math.Floor(y)
	x0 := int(x0f)
	y0 := int(y0f)
	x1 := x0 + 1
	y1 := y0 + 1

	noise.gradients[0] = noise.getGradient(x0, y0)
	noise.gradients[1] = noise.getGradient(x1, y0)
	noise.gradients[2] = noise.getGradient(x0, y1)
	noise.gradients[3] = noise.getGradient(x1, y1)

	noise.origins[0] = vector2{x0f + 0.0, y0f + 0.0}
	noise.origins[1] = vector2{x0f + 1.0, y0f + 0.0}
	noise.origins[2] = vector2{x0f + 0.0, y0f + 1.0}
	noise.origins[3] = vector2{x0f + 1.0, y0f + 1.0}
}

func (noise *Noise2D) Get2D(x, y float64) float64 {
	p := vector2{x, y}
	noise.getGradients(x, y)
	v0 := gradient(noise.origins[0], noise.gradients[0], p)
	v1 := gradient(noise.origins[1], noise.gradients[1], p)
	v2 := gradient(noise.origins[2], noise.gradients[2], p)
	v3 := gradient(noise.origins[3], noise.gradients[3], p)
	fx := smooth(x - noise.origins[0].x)
	vx0 := lerp(v0, v1, fx)
	vx1 := lerp(v2, v3, fx)
	fy := smooth(y - noise.origins[0].y)
	return lerp(vx0, vx1, fy)
}
