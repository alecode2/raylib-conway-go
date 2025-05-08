package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Panel struct {
	UIBase
	Color rl.Color
}

func NewPanel(color rl.Color) *Panel {
	return &Panel{
		UIBase: NewUIBase(),
		Color:  color,
	}
}

func (p *Panel) GetUIBase() *UIBase {
	return &p.UIBase
}

func (p *Panel) Draw() {
	rl.DrawRectangleRec(p.Bounds, p.Color)
}

func (p *Panel) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, p.Bounds)
}
