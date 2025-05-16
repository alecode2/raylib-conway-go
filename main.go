package main

import (
	"conway/assets"
	"conway/event"
	"conway/game"
	"conway/render"
	"conway/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	state := game.GameState{
		ScreenWidth:   int32(1280),
		ScreenHeight:  int32(720),
		FPS:           int32(60),
		SimFPS:        int32(5),
		CellSize:      int32(16),
		SelectedColor: rl.Red,
		IsMenuOpen:    true,
	}

	state.Step = 1 / float32(state.SimFPS)
	state.BoardA = game.NewBoard(int(state.ScreenWidth)/int(state.CellSize), int(state.ScreenHeight)/int(state.CellSize))
	state.BoardB = game.NewBoard(int(state.ScreenWidth)/int(state.CellSize), int(state.ScreenHeight)/int(state.CellSize))
	state.Current = state.BoardA
	state.Next = state.BoardB

	settings := game.Settings{
		FadeOpacity: false,
		FadeColor:   false,
		FadeLength:  5,
		GridColor:   rl.Gray,
	}

	// TODO: Delete this when we implement something more interesting
	game.AddShapes(state.Current)

	rl.InitWindow(state.ScreenWidth, state.ScreenHeight, "raylib game of life v0.0.6")

	rl.SetTargetFPS(state.FPS)

	//Initializing the Global Event Bus
	bus := event.NewEventBus()

	//Initializing the UI Tree
	root, registry := render.InitUI(state, settings, bus)

	//Subscribe game logic to events
	bus.Subscribe("toggle_pause", func(e event.Event) {
		state.IsPaused = !state.IsPaused
		fmt.Printf("Toggled Pause from Event Bus\n")
		if pausePanel, ok := registry["GREETING_LABEL"]; ok {
			pausePanel.GetUIBase().Visible = state.IsPaused
		}
	})

	// Setup clean termination function for proper resource cleanup
	cleanupFunc := func() {
		// Unload all textures
		assets.UnloadAllTextures()

		// Close the window
		rl.CloseWindow()
	}

	// Defer clean up to end of routine
	defer cleanupFunc()

	// Main rendering loop
	for !rl.WindowShouldClose() {
		//Reading inputs
		if rl.IsKeyPressed(rl.KeySpace) {
			bus.Emit(event.Event{Name: "toggle_pause"})
		}

		if rl.IsKeyPressed(rl.KeyRight) {
			state.StepForward = true
			fmt.Printf("hit rightArrow, requested Forwards Step=%v\n", state.StepForward)
		}

		if rl.IsKeyPressed(rl.KeyR) {
			state.ResetBoard()
		}

		mouse := rl.GetMousePosition()

		//Game Logic
		if state.Lapsed >= state.Step && !state.IsPaused {
			game.ConwayStep(&state)
			state.SwapBoards()
			state.Lapsed = 0
		}

		if state.IsPaused && state.StepForward {
			game.ConwayStep(&state)
			state.SwapBoards()
			state.StepForward = false
		}

		state.Lapsed += rl.GetFrameTime()

		//Handling UI Updates
		ui.RefreshUIEventList(root)
		ui.HandleUIHover(mouse)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			ui.HandleUIPress(mouse)
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			if !ui.HandleUIClick(mouse) {
				gridX := int32(mouse.X) / int32(state.CellSize)
				gridY := int32(mouse.Y) / int32(state.CellSize)
				state.ToggleCell(gridX, gridY)
			}

		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		render.DrawGrid(state.CellSize, state.Current, settings)
		render.DrawBoard(state.CellSize, state.Current, settings)

		ui.Size(root)
		ui.Position(root)
		ui.Draw(root)

		rl.EndDrawing()
	}

}
