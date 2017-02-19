package triangulate

type Polygon struct {
	Exterior  Ring
	Interiors []Ring
}

func (p Polygon) Triangulate() []Triangle {
	var result []Triangle
	ring := p.Exterior
	for len(ring) >= 3 {
		bestIndex := -1
		var bestScore float64
		for i := range ring {
			if ring.Ear(i) {
				score := ring.TriangleAt(i).MinAngle()
				if bestIndex < 0 || score > bestScore {
					bestIndex = i
					bestScore = score
				}
			}
		}
		if bestIndex < 0 {
			break
		}
		result = append(result, ring.TriangleAt(bestIndex))
		ring.Remove(bestIndex)
	}
	return result
}
