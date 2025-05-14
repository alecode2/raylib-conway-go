package main

import (
	"conway/assets"
	"conway/render"
	"conway/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "UI Test")
	rl.SetTargetFPS(60)

	root := render.InitUI()

	// Main rendering loop
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

		ui.Size(root)
		ui.Position(root)
		ui.Draw(root)

		rl.EndDrawing()
	}
	assets.UnloadAllTextures()
	rl.CloseWindow()
}
