package gographics

import (
	"image"
)

func DrawRect(img *image.Paletted, vec1, vec2 [2]int, col uint8) {
	DrawLine(img, [2]int{vec1[0], vec1[1]}, [2]int{vec2[0], vec1[1]}, col) // top line
	DrawLine(img, [2]int{vec1[0], vec2[1]}, [2]int{vec2[0], vec2[1]}, col) // bottom line
	DrawLine(img, [2]int{vec1[0], vec1[1]}, [2]int{vec1[0], vec2[1]}, col) // left line
	DrawLine(img, [2]int{vec2[0], vec1[1]}, [2]int{vec2[0], vec2[1]}, col) // right
}

func DrawFilledRect(img *image.Paletted, vec1, vec2 [2]int, col uint8) {
	// swap x1 and x2 if x1 is larger than x2
	if vec1[0] > vec2[0] {
		x := vec1[0]
		vec1[0] = vec2[0]
		vec2[0] = x
	}

	// swap y1 and y2 if y1 is larger than y2
	if vec1[1] > vec2[1] {
		y := vec1[1]
		vec1[1] = vec2[1]
		vec2[1] = y
	}

	for i := vec1[0]; i < vec2[0]; i++ {
		for k := vec1[1]; k < vec2[1]; k++ {
			img.SetColorIndex(i, k, col)
		}
	}
}
