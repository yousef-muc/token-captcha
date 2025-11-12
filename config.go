package tokencaptcha

import (
	"image/color"
	"time"
)

// FontConfig defines the font settings used when rendering captcha images.
// It allows either selecting a built-in font by name or providing a custom
// TrueType font as raw byte data. The Size and DPI fields control text scaling
// and resolution during rendering.
type FontConfig struct {
	Name string  // name of the built-in font (e.g. "noto-sans" or "jetbrains-mono")
	TTF  []byte  // optional raw TTF font data for custom fonts
	Size float64 // font size in points
	DPI  float64 // dots per inch used for text rendering
}

// Config defines all configurable parameters for the TokenCaptcha service.
// It controls how tokens are generated, verified, and optionally rendered as images.
// Missing or zero values are automatically normalized to defaults when creating
// a new Service instance using New(Config).
type Config struct {
	Secret        []byte        // secret key used for HMAC token signing
	Length        int           // number of characters in the captcha answer
	Expiry        time.Duration // lifetime of a captcha token before it expires
	Image         bool          // whether to generate a PNG image for the captcha
	Width         uint16        // image width in pixels
	Height        uint16        // image height in pixels
	Noise         uint16        // number of random noise lines drawn on the image
	CaseSensitive bool          // whether captcha answers are case-sensitive
	AllowActions  []string      // optional list of allowed action identifiers
	FG            color.Color   // foreground color for text
	BG            color.Color   // background color of the captcha image
	Font          FontConfig    // font configuration for text rendering
}
