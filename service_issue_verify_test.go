package tokencaptcha

import (
	"testing"
	"time"
)

func Test_IssueAndVerifyCaptcha(t *testing.T) {

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

	res, err := c.IssueCaptcha("signup")

	t.Run("Issue Captcha", func(t *testing.T) {
		if err != nil {
			t.Errorf("IssueCaptcha error: %v", err)
		}

		if res.TokenB64 == "" {
			t.Errorf("expected TokenB64")
		}

		if c.cfg.Image && res.PNGB64 == "" {
			t.Errorf("expected PNGB64 when Image=true")
		}

		_, err := decodeToken(res.TokenB64)
		if err != nil {
			t.Errorf("decodeToken error: %v", err)
		}

		if ok := c.Verify(res.TokenB64, "WRONG", "signup"); ok {
			t.Errorf("verify should fail with wrong answer")
		}
	})

}
