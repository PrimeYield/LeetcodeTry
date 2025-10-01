//3516. Find Closest Person

package math

func findClosest(x int, y int, z int) int {
	xToZ := 0
	yToZ := 0
	if z >= x {
		xToZ = z - x
	} else if z < x {
		xToZ = x - z
	}
	if z >= y {
		yToZ = z - y
	} else if z < y {
		yToZ = y - z
	}
	if xToZ > yToZ {
		return 2
	}
	if xToZ == yToZ {
		return 0
	}
	return 1
}
