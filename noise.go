package tokencaptcha

import (
	"image"
	"image/color"
	"math/rand/v2"
)

// addNoise draws a number of random colored lines on the given RGBA image.
// The noise is used to make captcha images harder to parse for automated systems.
// The number of lines is defined by n; if n is zero or negative, no noise is added.
func addNoise(img *image.RGBA, n int) {
	if n <= 0 {
		return
	}
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	for i := 0; i < n; i++ {
		c := color.RGBA{
			uint8(rand.IntN(256)),
			uint8(rand.IntN(256)),
			uint8(rand.IntN(256)),
			255,
		}
		x1 := rand.IntN(w)
		y1 := rand.IntN(h)
		x2 := rand.IntN(w)
		y2 := rand.IntN(h)
		drawLine(img, x1, y1, x2, y2, c)
	}
}

// drawLine renders a straight line on the image using the Bresenham algorithm (https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm).
// It ensures that all pixels of the line remain inside the image bounds.
func drawLine(img *image.RGBA, x0, y0, x1, y1 int, col color.Color) {
	dx := abs(x1 - x0)
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	dy := -abs(y1 - y0)
	sy := -1
	if y0 < y1 {
		sy = 1
	}
	err := dx + dy
	for {
		if image.Pt(x0, y0).In(img.Bounds()) {
			img.Set(x0, y0, col)
		}
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
