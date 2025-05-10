package main

import (
	"conway/ui"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "UI Test")
	rl.SetTargetFPS(60)

	// Create root panel (full screen)
	root := ui.NewPanel(rl.Gray)
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
	// Panel 1 (Fit Size)
	panel1 := ui.NewPanel(rl.Red)
	panel1_b := panel1.GetUIBase()
	panel1.ID = "Panel 1"
	panel1_b.WidthSizing = ui.SizingFit
	panel1_b.HeightSizing = ui.SizingFit
	panel1_b.PaddingTop = 16
	panel1_b.PaddingBottom = 16
	panel1_b.PaddingLeft = 16
	panel1_b.PaddingRight = 16
	panel1.EventHandlers[ui.EventClick] = func(evt ui.UIEvent) {
		fmt.Println("Clicked Panel 1")
	}
	ui.AddChild(root, panel1)

	// Sub Panel 1 (Fixed Size)
	subpanel1 := ui.NewPanel(rl.Green)
	subpanel1_b := subpanel1.GetUIBase()
	subpanel1.ID = "Subpanel"
	subpanel1_b.Width = 400
	subpanel1_b.Height = 300
	subpanel1_b.WidthSizing = ui.SizingFixed
	subpanel1_b.HeightSizing = ui.SizingFixed
	subpanel1.EventHandlers[ui.EventClick] = func(evt ui.UIEvent) {
		fmt.Println("Clicked SubPanel 1")
	}
	ui.AddChild(panel1, subpanel1)

	// Panel 2 (Grow Horizontal, Fixed Vertical)
	panel2 := ui.NewPanel(rl.Blue)
	panel2_b := panel2.GetUIBase()
	panel2_b.ID = "PANEL 2"
	panel2_b.MinWidth = 50
	panel2_b.MinHeight = 100
	panel2_b.WidthSizing = ui.SizingGrow
	panel2_b.HeightSizing = ui.SizingGrow
	panel2.EventHandlers[ui.EventPress] = func(evt ui.UIEvent) {
		fmt.Println("Pressed Panel 2")
	}
	ui.AddChild(root, panel2)

	ui.Size(root)
	ui.Position(root)

	fmt.Printf("Layout computed\n")

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
