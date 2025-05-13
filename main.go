package main

import (
	"conway/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "UI Test")
	rl.SetTargetFPS(60)

	// Create root panel (full screen)
	root := ui.NewPanel(rl.LightGray)
	root_b := root.GetUIBase()
	root_b.ID = "ROOT"
	root_b.Direction = ui.Horizontal
	root_b.Bounds.X = 0
	root_b.Bounds.Y = 0
	root_b.Width = 800
	root_b.Height = 600
	root_b.PaddingTop = 16
	root_b.PaddingBottom = 16
	root_b.PaddingLeft = 16
	root_b.PaddingRight = 16
	root_b.WidthSizing = ui.SizingFixed
	root_b.HeightSizing = ui.SizingFixed
	root_b.MainAlign = ui.AlignCenter
	root_b.CrossAlign = ui.CrossAlignCenter

	// 2) A label
	label := ui.NewLabel("Hello, Immediate-Mode GUI!", rl.GetFontDefault(), 24, rl.Maroon, ui.AlignTextCenter, false)
	lb := label.GetUIBase()
	lb.ID = "GREETING_LABEL"
	lb.WidthSizing = ui.SizingFit
	lb.HeightSizing = ui.SizingFit

	ui.AddChild(root, label)

	ui.Size(root)
	ui.Position(root)

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

		ui.Draw(root)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
