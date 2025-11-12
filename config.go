package tokencaptcha

import (
	"image/color"
	"time"
)

type FontConfig struct {
	Name string
	TTF  []byte
	Size float64
	DPI  float64
}

type Config struct {
	Secret        []byte
	Length        int
	Expiry        time.Duration
	Image         bool
	Width         uint16
	Height        uint16
	Noise         uint16
	CaseSensitive bool
	AllowActions  []string
	FG            color.Color
	BG            color.Color
	Font          FontConfig
}
