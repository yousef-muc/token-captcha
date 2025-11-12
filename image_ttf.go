package tokencaptcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// renderPNGBase64TTF renders the provided text into a PNG image using a TrueType font
// and returns the result as a Base64-encoded string.
// The image is generated entirely in memory using the standard library and the golang.org/x/image package.
// It applies background color, foreground color, random noise lines for obfuscation,
// and centers the text horizontally and vertically inside the image.
// The resulting Base64-encoded string can be directly used in data URLs such as
// "data:image/png;base64,<string>" for frontend display.
func renderPNGBase64TTF(text string, cfg Config) (string, error) {
	w := int(cfg.Width)
	h := int(cfg.Height)

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	bg := pickColor(cfg.BG, color.RGBA{245, 245, 245, 255})
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)

	addNoise(img, int(cfg.Noise))

	face, err := loadFace(cfg.Font)
	if err != nil {
		return "", err
	}
	defer face.Close()

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(pickColor(cfg.FG, color.Black)),
		Face: face,
	}

	txt := text
	adv := drawer.MeasureString(txt)
	textW := int(adv.Round())

	met := face.Metrics()
	ascent := met.Ascent.Ceil()
	descent := met.Descent.Ceil()
	textH := ascent + descent

	x := (w - textW) / 2
	y := (h+textH)/2 - descent

	drawer.Dot = fixed.P(x, y)
	drawer.DrawString(txt)

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// pickColor returns the provided color if not nil; otherwise it returns the given fallback color.
// This helper ensures that color configuration always produces a valid image color value.
func pickColor(c color.Color, fallback color.Color) color.Color {
	if c == nil {
		return fallback
	}
	return c
}

// newFaceFromFont creates a new scalable font.Face instance from a parsed TrueType font.
// The size and DPI values determine the final rendering scale.
// This helper provides fine-grained control over text rendering when using custom font sources.
func newFaceFromFont(f *opentype.Font, size, dpi float64) (font.Face, error) {
	return opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
