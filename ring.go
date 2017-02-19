package triangulate

import "math"

type Ring []Point

func (r Ring) At(i int) Point {
	n := len(r)
	return r[(i+n)%n]
}

func (r *Ring) Remove(i int) {
	*r = append((*r)[:i], (*r)[i+1:]...)
}

func (r Ring) Ear(i int) bool {
	t := r.TriangleAt(i)
	a := t.A.Sub(t.B).Atan2() - t.C.Sub(t.B).Atan2()
	if a <= 0 || a >= math.Pi {
		return false
	}
	for j := 0; j < len(r)-3; j++ {
		p := r.At(i + 2 + j)
		if t.Contains(p) {
			return false
		}
	}
	return true
}

func (r Ring) TriangleAt(i int) Triangle {
	return Triangle{r.At(i - 1), r.At(i), r.At(i + 1)}
}
