package input

import (
	"conway/event"
	"conway/game"
	state "conway/game"
	ui "conway/ui"
	cmp "conway/ui/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	currentMode       = ModeGameplay
	activeTextFieldID string
	mouse             rl.Vector2
)

func SetInputMode(inputMode InputMode) {
	currentMode = inputMode
}

func InitInput(bus *event.EventBus) {
	bus.Subscribe("focus_hex_input", func(e event.Event) {
		currentMode = ModeTextInput
		activeTextFieldID = "HEX_INPUT"
	})
}

func HandleInput(game *state.GameState, bus *event.EventBus) {
	mouse = rl.GetMousePosition()

	switch currentMode {
	case ModeGameplay:
		handleGameplayInput(game, bus)
	case ModeTextInput:
		handleTextInput(game, bus)
	}
}

func handleGameplayInput(state *state.GameState, bus *event.EventBus) {
	mouse := rl.GetMousePosition()

	// Refresh the UI event list with the current tree
	ui.RefreshUIEventList(state.UIRoot)

	// --- 1. Handle mouse press
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if ui.HandleUIPress(mouse) {
			return // UI consumed the press
		}
	}

	// --- 2. Handle mouse release
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		if ui.HandleUIRelease(mouse) {
			return // UI consumed the release and/or click
		}

		// UI didn't consume, treat as gameplay click
		gridX := int32(mouse.X) / int32(state.CellSize)
		gridY := int32(mouse.Y) / int32(state.CellSize)
		state.ToggleCell(gridX, gridY)
	}

	// --- 3. Handle hover updates
	ui.HandleUIHover(mouse)

	// --- 4. Handle keyboard input (always available in gameplay)
	for key, action := range gameplayActionBindings {
		if rl.IsKeyPressed(key) {
			bus.Emit(event.Event{Name: action})
		}
	}

	for key, fn := range gameplayFuncBindings {
		if rl.IsKeyPressed(key) {
			fn(state)
		}
	}
}

func handleTextInput(state *game.GameState, bus *event.EventBus) {

	fieldElem, ok := state.UIMap[activeTextFieldID]
	if !ok {
		return
	}
	field, ok := fieldElem.(*cmp.InputField)
	if !ok {
		return
	}

	// Blur on pressing enter
	if rl.IsKeyPressed(rl.KeyEnter) {
		bus.Emit(event.Event{Name: "hex_input_submit", Data: field.Value})
		currentMode = ModeGameplay
		activeTextFieldID = ""
		return
	}

	// Blur on click outside
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if !rl.CheckCollisionPointRec(mouse, field.Bounds) {
			currentMode = ModeGameplay
			activeTextFieldID = ""
			return
		}
	}

	// Backspace erases up to the first #
	if rl.IsKeyPressed(rl.KeyBackspace) && len(field.Value) > 1 {
		field.Value = field.Value[:len(field.Value)-1]
		field.Label.Text = field.Value
	}

	// If the string size is 7 the hex color code is complete, we ignore further characters
	if len(field.Value) == 7 {
		return
	}

	// Hex input
	char := rl.GetCharPressed()
	for char > 0 {
		r := rune(char)
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F') || r == '#' {
			field.Value += string(r)
			field.Label.Text = field.Value
		}
		char = rl.GetCharPressed()
	}
}
