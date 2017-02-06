package math

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
