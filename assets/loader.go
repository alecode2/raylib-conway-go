package assets

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var textures = make(map[string]rl.Texture2D)
var fonts = make(map[string]rl.Font)

// LoadTexture loads a texture if it hasn't been loaded already
func LoadTexture(path string) rl.Texture2D {
	if tex, exists := textures[path]; exists {
		return tex
	}

	tex := rl.LoadTexture(path)
	if tex.ID == 0 {
		panic(fmt.Sprintf("Failed to load texture: %s", path))
	}

	textures[path] = tex
	return tex
}

// GetTexture returns a previously loaded texture, panics if not found
func GetTexture(path string) rl.Texture2D {
	tex, ok := textures[path]
	if !ok {
		panic(fmt.Sprintf("Texture not loaded: %s", path))
	}
	return tex
}

// UnloadAllTextures unloads all loaded textures (call on shutdown)
func UnloadAllTextures() {
	for path, tex := range textures {
		fmt.Printf("Unloading: %s\n", path)
		rl.UnloadTexture(tex)
	}
	textures = make(map[string]rl.Texture2D) // Clear map
}

func LoadFont(path string, baseSize int32) rl.Font {
	if font, exists := fonts[path]; exists {
		return font
	}
	font := rl.LoadFontEx(path, baseSize, nil)
	rl.GenTextureMipmaps(&font.Texture)

	if font.Texture.ID == 0 {
		panic(fmt.Sprintf("Failed to load font: %s", path))
	}

	fonts[path] = font
	return font
}

// GetTexture returns a previously loaded texture, panics if not found
func GetFont(path string) rl.Font {
	font, ok := fonts[path]
	if !ok {
		panic(fmt.Sprintf("Font not loaded: %s", path))
	}
	return font
}

func UnloadAllFonts() {
	for path, font := range fonts {
		fmt.Printf("Unloading: %s\n", path)
		rl.UnloadFont(font)
	}
	fonts = make(map[string]rl.Font) // Clear map
}
