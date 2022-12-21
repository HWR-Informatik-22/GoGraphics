package gographics

import (
	"image"
)

func DrawLine(img *image.Paletted, vec1, vec2 [2]int, col uint8) {
	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if vec1[0] > vec2[0] {
		vec1[0], vec1[1], vec2[0], vec2[1] = vec2[0], vec2[1], vec1[0], vec1[1]
	}

	dx, dy = vec2[0]-vec1[0], vec2[1]-vec1[1]
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case vec1[0] == vec2[0] && vec1[1] == vec2[1]:
		img.SetColorIndex(vec1[0], vec1[1], col)

	// Is line an horizontal ?
	case vec1[1] == vec2[1]:
		for ; dx != 0; dx-- {
			img.SetColorIndex(vec1[0], vec1[1], col)
			vec1[0]++
		}
		img.SetColorIndex(vec1[0], vec1[1], col)

	// Is line a vertical ?
	case vec1[0] == vec2[0]:
		if vec1[1] > vec2[1] {
			vec1[1], vec2[1] = vec2[1], vec1[1]
		}
		for ; dy != 0; dy-- {
			img.SetColorIndex(vec1[0], vec1[1], col)
			vec1[1]++
		}
		img.SetColorIndex(vec1[0], vec1[1], col)

	// Is line a diagonal ?
	case dx == dy:
		if vec1[1] < vec2[1] {
			for ; dx != 0; dx-- {
				img.SetColorIndex(vec1[0], vec1[1], col)
				vec1[0]++
				vec1[1]++
			}
		} else {
			for ; dx != 0; dx-- {
				img.SetColorIndex(vec1[0], vec1[1], col)
				vec1[0]++
				vec1[1]--
			}
		}
		img.SetColorIndex(vec1[0], vec1[1], col)

	// wider than high ?
	case dx > dy:
		if vec1[1] < vec2[1] {
			// BresenhamDxXRYD(img, vec1[0], vec1[1], vec2[0], vec2[1], col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.SetColorIndex(vec1[0], vec1[1], col)
				vec1[0]++
				e -= dy
				if e < 0 {
					vec1[1]++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, vec1[0], vec1[1], vec2[0], vec2[1], col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.SetColorIndex(vec1[0], vec1[1], col)
				vec1[0]++
				e -= dy
				if e < 0 {
					vec1[1]--
					e += slope
				}
			}
		}
		img.SetColorIndex(vec2[0], vec2[1], col)

	// higher than wide.
	default:
		if vec1[1] < vec2[1] {
			// BresenhamDyXRYD(img, vec1[0], vec1[1], vec2[0], vec2[1], col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.SetColorIndex(vec1[0], vec1[1], col)
				vec1[1]++
				e -= dx
				if e < 0 {
					vec1[0]++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, vec1[0], vec1[1], vec2[0], vec2[1], col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.SetColorIndex(vec1[0], vec1[1], col)
				vec1[1]--
				e -= dx
				if e < 0 {
					vec1[0]++
					e += slope
				}
			}
		}
		img.SetColorIndex(vec2[0], vec2[1], col)
	}
}
