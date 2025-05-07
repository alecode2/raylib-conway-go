package render

import (
	"conway/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawGrid(tileSize int32, board game.Board, settings game.Settings) {
	for x, column := range board {
		for y := range column {
			rect := rl.Rectangle{
				X:      float32(int32(x) * tileSize),
				Y:      float32(int32(y) * tileSize),
				Width:  float32(tileSize),
				Height: float32(tileSize),
			}

			rl.DrawRectangleLinesEx(rect, float32(.5), settings.GridColor)
		}
	}
}

func DrawBoard(tileSize int32, board game.Board, settings game.Settings) {
	for x, column := range board {
		for y, tile := range column {
			tileColor := GetTileColor(tile, settings)

			rect := rl.Rectangle{
				X:      float32(int32(x) * tileSize),
				Y:      float32(int32(y) * tileSize),
				Width:  float32(tileSize),
				Height: float32(tileSize),
			}
			rl.DrawRectangleRec(rect, tileColor)
		}
	}
}

func GetTileColor(tile game.Tile, settings game.Settings) rl.Color {
	color := tile.Color
	// Only apply fading if the tile is dead
	if !tile.Alive {
		// Apply opacity fade if enabled
		if settings.FadeOpacity && tile.Age < uint8(settings.FadeLength) {
			// Apply opacity fade based on the tile's age
			value := rl.Remap(float32(tile.Age), 0, float32(settings.FadeLength), 255, 0)

			color = rl.Color{R: color.R, G: color.G, B: color.B, A: uint8(value)}
		} else {
			color.A = 0
		}
	}

	return color
}
