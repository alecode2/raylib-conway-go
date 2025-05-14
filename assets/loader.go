package assets

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var textures = make(map[string]rl.Texture2D)

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
