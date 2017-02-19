package triangulate

import "math"

type Point struct {
	X, Y float64
}

func (a Point) Normalize() Point {
	r := 1 / math.Sqrt(a.X*a.X+a.Y*a.Y)
	return Point{a.X * r, a.Y * r}
}

func (a Point) Dot(b Point) float64 {
	return a.X*b.X + a.Y*b.Y
}

func (a Point) Sub(b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y}
}

func (a Point) Atan2() float64 {
	return math.Atan2(a.Y, a.X)
}
