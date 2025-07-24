package main

import (
	"conway/assets"
	"conway/event"
	"conway/game"
	"conway/input"
	"conway/render"
	"conway/ui"

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

	//Initializing the tools
	game.InitToolBox(&state, &settings, bus)

	//Initializing the UI Tree
	_, _ = game.InitUI(&state, &settings, bus)

	//Initializing the Input Event listeners
	input.InitInput(bus)

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
		delta := rl.GetFrameTime()

		input.HandleInput(&state, bus)

		//Simulation
		if state.Lapsed >= state.Step && !state.IsPaused {
			game.ConwayStep(&state)
			state.SwapBoards()
			state.Lapsed = 0
		}

		if state.IsPaused && state.StepForward {
			game.ConwayStep(&state)
			state.SwapBoards()
			state.Lapsed = 0
			state.StepForward = false
		}

		state.Lapsed += delta

		//UI Updates
		ui.Size(state.UIRoot)
		ui.Position(state.UIRoot)
		ui.AdvanceAnimations(state.UIRoot, delta)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		render.DrawGrid(state.CellSize, state.Current, settings)
		render.DrawBoard(state.CellSize, state.Current, settings)

		ui.Draw(state.UIRoot)

		rl.EndDrawing()
	}

}
