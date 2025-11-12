package tokencaptcha

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"strconv"
)

// tokenPayload defines the internal structure of a captcha token.
// Each token contains a random nonce, an expiry timestamp, the related action name,
// and an HMAC-SHA256 signature to ensure integrity and authenticity.
type tokenPayload struct {
	C string `json:"c"` // nonce value used as random salt
	E int64  `json:"e"` // expiry time as UNIX timestamp
	A string `json:"a"` // action identifier associated with the captcha
	M string `json:"m"` // HMAC-SHA256 checksum in hexadecimal format
}

// encodeToken serializes a tokenPayload to JSON and encodes it using Base64URL without padding.
// The resulting string is safe to include in URLs and API responses.
func encodeToken(t tokenPayload) (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data), nil
}

// decodeToken decodes a Base64URL-encoded JSON string back into a tokenPayload structure.
// Returns an error if decoding or JSON unmarshalling fails.
func decodeToken(b64url string) (tokenPayload, error) {
	var t tokenPayload
	data, err := base64.RawURLEncoding.DecodeString(b64url)
	if err != nil {
		return t, err
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return t, err
	}
	return t, nil
}

// mac calculates the HMAC-SHA256 checksum for a given captcha answer, nonce, expiry, and action.
// The message format is "answer|nonce|expiry|action" and the result is returned as a hexadecimal string.
// This function ensures that token validation can be performed statelessly on any server.
func mac(secret []byte, answer, c string, e int64, a string) string {
	h := hmac.New(sha256.New, secret)
	io.WriteString(h, answer+"|"+c+"|"+strconv.FormatInt(e, 10)+"|"+a)
	return hex.EncodeToString(h.Sum(nil))
}
