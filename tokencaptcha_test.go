package tokencaptcha

import (
	"reflect"
	"testing"
	"time"
)

func Test_New(t *testing.T) {

	tests := []struct {
		name         string
		expectedType string
	}{
		{"Typecheck", "*tokencaptcha.Service"},
	}

	c := New(Config{
		Secret:        []byte("ITS-A-TEST"),
		Length:        12,
		Expiry:        5 * time.Minute,
		Width:         10,
		Height:        20,
		Image:         true,
		CaseSensitive: true,
		Font: FontConfig{
			Name: "jetbrains-mono",
			Size: 12,
			DPI:  30,
		},
	})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := reflect.TypeOf(c).String(); got != "*tokencaptcha.Service" {
				t.Errorf("Expected type %s, but got %s\n", test.expectedType, reflect.TypeOf(c).String())
			}
		})
	}
}

func Test_Defaults(t *testing.T) {

	tests := []struct {
		name   string
		secret string
		length int
		expiry time.Duration
		width  uint16
		height uint16
		noise  uint16
		size   float64
		dpi    float64
	}{
		{"Value check", "MY-TOKEN-SECRET", 6, time.Minute * 2, 220, 70, 10, 28, 72},
	}

	c := New(Config{})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if string(c.cfg.Secret) != test.secret {
				t.Errorf("Expected secret %s, but got %s", test.secret, string(c.cfg.Secret))
			}

			if c.cfg.Length != test.length {
				t.Errorf("Expected length %d, but got %d", test.length, c.cfg.Length)
			}

			if c.cfg.Expiry.Minutes() != test.expiry.Minutes() {
				t.Errorf("Expected expiry %v, but got %v", test.expiry.Minutes(), c.cfg.Expiry.Minutes())
			}

			if c.cfg.Width != test.width {
				t.Errorf("Expected width %d, but got %d", test.width, c.cfg.Width)
			}

			if c.cfg.Height != test.height {
				t.Errorf("Expected height %d, but got %d", test.height, c.cfg.Height)
			}

			if c.cfg.Noise != test.noise {
				t.Errorf("Expected noise %d, but got %d", test.noise, c.cfg.Noise)
			}

			if uint16(c.cfg.Font.Size) != uint16(test.size) {
				t.Errorf("Expected font size %d, but got %d", uint16(test.size), uint16(c.cfg.Font.Size))
			}

			if uint16(c.cfg.Font.DPI) != uint16(test.dpi) {
				t.Errorf("Expected dpi %d, but got %d", uint16(test.dpi), uint16(c.cfg.Font.DPI))
			}
		})
	}
}
