package triangulate

import "math"

type Triangle struct {
	A, B, C Point
}

func (t Triangle) Contains(p Point) bool {
	const eps = 1e-9
	r := 1 / (-t.B.Y*t.C.X + t.A.Y*(-t.B.X+t.C.X) + t.A.X*(t.B.Y-t.C.Y) + t.B.X*t.C.Y)
	u := (t.A.Y*t.C.X - t.A.X*t.C.Y + (t.C.Y-t.A.Y)*p.X + (t.A.X-t.C.X)*p.Y) * r
	v := (t.A.X*t.B.Y - t.A.Y*t.B.X + (t.A.Y-t.B.Y)*p.X + (t.B.X-t.A.X)*p.Y) * r
	if u < -eps || u > 1+eps || v < -eps || v > 1+eps || u+v > 1+eps {
		return false
	}
	return true
	// b1 := (t.A.X-t.B.X)*(p.Y-t.B.Y)-(t.A.Y-t.B.Y)*(p.X-t.B.X) < 1e-9
	// b2 := (t.B.X-t.C.X)*(p.Y-t.C.Y)-(t.B.Y-t.C.Y)*(p.X-t.C.X) < 1e-9
	// b3 := (t.C.X-t.A.X)*(p.Y-t.A.Y)-(t.C.Y-t.A.Y)*(p.X-t.A.X) < 1e-9
	// return b1 == b2 && b2 == b3
}

func (t Triangle) MinAngle() float64 {
	a := math.Acos(t.C.Sub(t.A).Normalize().Dot(t.B.Sub(t.A).Normalize()))
	b := math.Acos(t.A.Sub(t.B).Normalize().Dot(t.C.Sub(t.B).Normalize()))
	c := math.Acos(t.B.Sub(t.C).Normalize().Dot(t.A.Sub(t.C).Normalize()))
	return math.Min(a, math.Min(b, c))
}
