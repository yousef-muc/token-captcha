package tokencaptcha

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// charset defines the allowed set of characters used for captcha answers.
// The characters I, O, 1, and 0 are intentionally excluded because they
// are visually ambiguous and may cause confusion for users.
var charset = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ23456789")

// randomAnswer generates a random captcha string of length n using the predefined charset.
// If n is less than or equal to zero, a default length of six characters is used.
// The result is returned as an uppercase string. The randomness is provided by
// crypto/rand to ensure cryptographic security.
func randomAnswer(n int) (string, error) {
	if n <= 0 {
		n = 6
	}
	rb := make([]byte, n)
	if _, err := rand.Read(rb); err != nil {
		return "", err
	}
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = charset[int(rb[i])%len(charset)]
	}
	return strings.ToUpper(string(runes)), nil
}

// mustNonce generates a cryptographically secure random nonce of length n bytes
// and returns it as a Base64URL-encoded string without padding.
// This function ignores potential read errors, as nonce generation failures
// are highly unlikely in typical system environments.
func mustNonce(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}
