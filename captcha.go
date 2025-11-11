package tokencaptcha

import "fmt"

type TokenCaptcha struct {
	length uint32
}

func New(l uint32) *TokenCaptcha {
	return &TokenCaptcha{
		length: l,
	}
}

func (t *TokenCaptcha) Print() {
	fmt.Printf("length: %d\n", t.length)
}
