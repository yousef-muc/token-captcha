package tokencaptcha

import (
	"crypto/hmac"
	"strings"
	"time"
)

// IssueResult represents the response of a generated captcha.
// TokenB64 contains the Base64URL-encoded JSON token payload.
// PNGB64 contains the optional captcha image in Base64 PNG format if image generation is enabled.
type IssueResult struct {
	TokenB64 string // Base64URL JSON tokenPayload
	PNGB64   string // optional field, only set if cfg.Image is true
}

// IssueCaptcha creates a new stateless captcha token and optionally an image.
// It generates a random answer, computes the HMAC signature, encodes the payload,
// and returns the token along with the image if configured.
// The returned captcha is completely stateless and can be verified later
// using the Verify method without any stored session data.
func (s *Service) IssueCaptcha(action string) (IssueResult, error) {
	ans, err := randomAnswer(s.cfg.Length)
	if err != nil {
		return IssueResult{}, nil
	}

	nonce := mustNonce(16)
	exp := time.Now().Add(s.cfg.Expiry).Unix()

	macHex := mac(s.cfg.Secret, normalizeAnswer(ans, s.cfg.CaseSensitive), nonce, exp, action)
	p := tokenPayload{C: nonce, E: exp, A: action, M: macHex}

	tb64, err := encodeToken(p)
	if err != nil {
		return IssueResult{}, nil
	}

	res := IssueResult{TokenB64: tb64}

	if s.cfg.Image {
		pngB64, err := renderPNGBase64TTF(ans, s.cfg)
		if err != nil {
			return IssueResult{}, nil
		}
		res.PNGB64 = pngB64
	}

	return res, nil
}

// Verify validates a user provided captcha answer against a given token.
// It decodes the token, checks expiration and action values,
// recalculates the expected HMAC signature, and performs a constant-time comparison.
// Returns true if the captcha is valid and false otherwise.
func (s *Service) Verify(tokenB64, userAnswer, expectedAction string) bool {
	t, err := decodeToken(tokenB64)
	if err != nil {
		return false
	}
	if time.Now().Unix() > t.E {
		return false
	}
	if expectedAction != "" && t.A != expectedAction {
		return false
	}
	if len(s.cfg.AllowActions) > 0 && !contains(s.cfg.AllowActions, t.A) {
		return false
	}

	expected := mac(s.cfg.Secret, normalizeAnswer(userAnswer, s.cfg.CaseSensitive), t.C, t.E, t.A)

	return hmac.Equal([]byte(expected), []byte(t.M))
}

// normalizeAnswer trims surrounding spaces and converts the string
// to lowercase if case sensitivity is disabled in the configuration.
func normalizeAnswer(s string, caseSensitive bool) string {
	x := strings.TrimSpace(s)
	if caseSensitive {
		return x
	}
	return strings.ToLower(x)
}

// contains checks whether the provided slice contains the given string value.
func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
