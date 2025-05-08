/*
package main

import (
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

	// TODO: Delete this when we implement something more interesting
	game.AddShapes(state.Current)

	settings := game.Settings{
		FadeOpacity: false,
		FadeColor:   false,
		FadeLength:  5,
		GridColor:   rl.Gray,
	}

	rl.InitWindow(state.ScreenWidth, state.ScreenHeight, "raylib game of life v0.0.6")
	defer rl.CloseWindow()

	rl.SetTargetFPS(state.FPS)

	//Initializing the Event Bus
	bus := event.NewEventBus()

	//Initializing the UIElements slice
	UITree, UIRegistry := render.GetUI(state, settings, bus)

	//Subscribe game logic to events
	bus.Subscribe("toggle_pause", func(e event.Event) {
		state.IsPaused = !state.IsPaused
		fmt.Printf("Toggled Pause from Event Bus\n")
		if pausePanel, ok := UIRegistry["pause_panel"]; ok {
			pausePanel.SetVisible(state.IsPaused)
		}
	})

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

		//Handling UI Updates
		ui.RefreshUIEventList(UITree)
		ui.HandleUIHover(mouse)

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			ui.HandleUIPress(mouse)
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			if !ui.HandleUIClick(mouse) {
				gridX := int32(mouse.X) / int32(state.CellSize)
				gridY := int32(mouse.Y) / int32(state.CellSize)
				state.ToggleCell(gridX, gridY)
			}

		}

		state.Lapsed += rl.GetFrameTime()

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

		//Drawing the creen
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		render.DrawGrid(state.CellSize, state.Current, settings)
		render.DrawBoard(state.CellSize, state.Current, settings)
		render.DrawUI(UITree)

		rl.EndDrawing()
	}
}
*/

package main

import (
	"conway/ui"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "UI Test")
	rl.SetTargetFPS(10)

	// Create root panel
	root := ui.NewPanel(rl.Gray)
	ui.SetBounds(root, rl.NewRectangle(0, 0, 800, 600))

	// Panel 1
	panel1 := ui.NewPanel(rl.Red)
	ui.SetBounds(panel1, rl.NewRectangle(100, 100, 200, 100))
	panel1.EventHandlers[ui.EventHover] = func(evt ui.UIEvent) {
		fmt.Println("Hovered Panel 1")
	}
	panel1.EventHandlers[ui.EventClick] = func(evt ui.UIEvent) {
		fmt.Println("Clicked Panel 1")
	}
	ui.AddChild(root, panel1)

	// Panel 2
	panel2 := ui.NewPanel(rl.Blue)
	ui.SetBounds(panel2, rl.NewRectangle(350, 100, 200, 100))
	panel2.EventHandlers[ui.EventHover] = func(evt ui.UIEvent) {
		fmt.Println("Hovered Panel 2")
	}
	panel2.EventHandlers[ui.EventClick] = func(evt ui.UIEvent) {
		fmt.Println("Clicked Panel 2")
	}
	ui.AddChild(root, panel2)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		mouse := rl.GetMousePosition()
		ui.RefreshUIEventList(root)
		ui.HandleUIHover(mouse)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			ui.HandleUIPress(mouse)
		}
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			ui.HandleUIClick(mouse)
		}

		ui.Draw(root)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
