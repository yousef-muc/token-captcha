package tokencaptcha

import (
	"image/color"
	"time"
)

// Service represents the main captcha generator and verifier.
// It operates in a stateless manner, meaning that no data needs to be stored on the server.
// The service uses an HMAC signature to verify captcha tokens without maintaining sessions.
type Service struct {
	cfg Config
}

// New creates a new captcha service using the provided configuration.
// It automatically normalizes the configuration by filling in missing or invalid values
// with secure and reasonable defaults before returning the initialized Service instance.
func New(cfg Config) *Service {
	normalizeConfig(&cfg)
	return &Service{
		cfg: cfg,
	}
}

// normalizeConfig ensures that all configuration values are properly set.
// Missing values are replaced with secure defaults such as a default secret key,
// default image size, and default font parameters.
// Boolean fields (like Image or CaseSensitive) are not modified here and are expected
// to be set explicitly by the user according to the desired behavior.
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
	if c.Width == 0 {
		c.Width = 220
	}
	if c.Height == 0 {
		c.Height = 70
	}
	if c.Noise == 0 {
		c.Noise = 10
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
