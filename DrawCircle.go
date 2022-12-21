package gographics

import (
	"image"
)

// https://iq.opengenus.org/bresenhams-circle-drawing-algorithm/
func DrawCircle(img *image.Paletted, vec [2]int, radius int, col uint8) {
	x := 0
	y := radius
	d := 3 - 2*radius

	displayBresenhamCirlce(img, vec[0], vec[1], x, y, col)

	for y >= x {
		x++

		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}

		displayBresenhamCirlce(img, vec[0], vec[1], x, y, col)
	}
}

// https://stackoverflow.com/questions/1201200/fast-algorithm-for-drawing-filled-circles
func DrawFilledCirlce(img *image.Paletted, vec [2]int, radius int, col uint8) {
	x := 0
	y := radius
	d := 3 - 2*radius

	displayBresenhamCirlce(img, vec[0], vec[1], x, y, col)
	fillBresenhamCircle(img, vec[0], vec[1], x, y, col)

	for y >= x {
		x++

		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}

		displayBresenhamCirlce(img, vec[0], vec[1], x, y, col)
		fillBresenhamCircle(img, vec[0], vec[1], x, y, col)
	}
}

func displayBresenhamCirlce(img *image.Paletted, xc, yc, x, y int, col uint8) {
	img.SetColorIndex(xc+x, yc+y, col)
	img.SetColorIndex(xc-x, yc+y, col)
	img.SetColorIndex(xc+x, yc-y, col)
	img.SetColorIndex(xc-x, yc-y, col)
	img.SetColorIndex(xc+y, yc+x, col)
	img.SetColorIndex(xc-y, yc+x, col)
	img.SetColorIndex(xc+y, yc-x, col)
	img.SetColorIndex(xc-y, yc-x, col)
}

func fillBresenhamCircle(img *image.Paletted, xc, yc, x, y int, col uint8) {
	for i := xc - x; i <= xc+x; i++ {
		img.SetColorIndex(i, yc+y, col)
		img.SetColorIndex(i, yc-y, col)
	}

	for i := xc - y; i <= xc+y; i++ {
		img.SetColorIndex(i, yc+x, col)
		img.SetColorIndex(i, yc-x, col)
	}
}
