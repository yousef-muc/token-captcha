package tokencaptcha

import (
	"image/color"
	"time"
)

type Service struct {
	cfg Config
}

func New(cfg Config) *Service {
	normalizeConfig(&cfg)
	return &Service{
		cfg: cfg,
	}
}

func normalizeConfig(c *Config) {
	if c.Secret == nil {
		c.Secret = []byte("MY-TOKEN-SECRET")
	}
	if c.Length <= 0 {
		c.Length = 6
	}
	if c.Expiry <= 0 {
		c.Expiry = 2 * time.Minute
	}
	if !c.Image {
		c.Image = true
	}
	if c.Width == 0 {
		c.Width = 220
	}
	if c.Height == 0 {
		c.Height = 70
	}
	if c.Noise == 0 {
		c.Noise = 10
	}
	if !c.CaseSensitive {
		c.CaseSensitive = true
	}
	if c.FG == nil {
		c.FG = color.Black
	}
	if c.BG == nil {
		c.BG = color.RGBA{245, 245, 245, 255}
	}
	if c.Font.Size <= 0 {
		c.Font.Size = 28
	}
	if c.Font.DPI <= 0 {
		c.Font.DPI = 72
	}
}
