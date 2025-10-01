//812. Largest Triangle Area

package math

import "math"

func largestTriangleArea(points [][]int) float64 {
	var max float64
	for a := 0; a < len(points); a++ {
		for b := a + 1; b < len(points); b++ {
			for c := b + 1; c < len(points); c++ {
				abVector := []float64{float64(points[b][0]) - float64(points[a][0]), float64(points[b][1]) - float64(points[a][1])}
				acVector := []float64{float64(points[c][0]) - float64(points[a][0]), float64(points[c][1]) - float64(points[a][1])}
				area := (math.Abs(abVector[0]*acVector[1] - abVector[1]*acVector[0])) / 2
				if max < area {
					max = area
				}
			}
		}
	}
	return max
}
