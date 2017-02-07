package math

import "math"

type Rect struct {
	X, Y, W, H int
}

type Vector3 struct {
	X, Y, Z int
}

func Clamp(x, min, max int) int {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}

func DivCeil(a, b int) int {
	n := math.Ceil(float64(a) / float64(b))
	return int(n)
}

func DivFloor(a, b int) int {
	n := math.Floor(float64(a) / float64(b))
	return int(n)
}
