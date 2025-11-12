package tokencaptcha

import (
	"embed"
	"io/fs"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// fontFS embeds all TrueType font files from the assets/fonts directory.
// These fonts are used as built-in defaults when no external font is provided by the user.
//
//go:embed assets/fonts/*.ttf
var fontFS embed.FS

// Built-in font identifiers for easier configuration and selection.
const (
	fontDefault   = "default"
	fontNotoSans  = "noto-sans"
	fontJetBrains = "jetbrains-mono"
)

// builtInFonts maps internal font identifiers to their corresponding embedded font file paths.
// The default font is NotoSans-Regular.ttf, which provides broad Unicode coverage
// and good readability for captcha rendering.
var builtInFonts = map[string]string{
	fontDefault:   "assets/fonts/NotoSans-Regular.ttf",
	fontNotoSans:  "assets/fonts/NotoSans-Regular.ttf",
	fontJetBrains: "assets/fonts/JetBrainsMono-Regular.ttf",
}

// loadFace loads and initializes a font.Face based on the provided FontConfig.
// It supports three modes:
// 1. Direct TTF data (FontConfig.TTF) provided by the user.
// 2. Built-in font selection by name (FontConfig.Name) from the embedded assets.
// 3. Default fallback font if no valid font name or data is found.
//
// The returned font.Face can be used for text rendering with the golang.org/x/image/font package.
func loadFace(cfg FontConfig) (font.Face, error) {
	// 1. Direct byte data provided by the user
	if len(cfg.TTF) > 0 {
		f, err := opentype.Parse(cfg.TTF)
		if err != nil {
			return nil, err
		}
		return newFaceFromFont(f, cfg.Size, cfg.DPI)
	}

	// 2. Resolve font name to an embedded file path
	name := strings.ToLower(strings.TrimSpace(cfg.Name))
	if name == "" {
		name = fontDefault
	}
	path, ok := builtInFonts[name]
	if !ok {
		path = builtInFonts[fontDefault]
	}

	// 3. Load font bytes from embedded filesystem with fallback to default font
	data, err := fs.ReadFile(fontFS, path)
	if err != nil {
		data, err = fs.ReadFile(fontFS, builtInFonts[fontDefault])
		if err != nil {
			return nil, err
		}
	}

	f, err := opentype.Parse(data)
	if err != nil {
		return nil, err
	}
	return newFaceFromFont(f, cfg.Size, cfg.DPI)
}
