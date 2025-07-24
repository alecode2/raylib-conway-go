package input

import (
	"conway/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Map from keys to event names
var gameplayActionBindings = map[int32]string{
	rl.KeySpace: "toggle_pause",
}

// Map from keys to direct state functions
var gameplayFuncBindings = map[int32]func(*game.GameState){
	rl.KeyR: func(s *game.GameState) {
		s.ResetBoard()
	},
}
