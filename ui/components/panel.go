package cmp

import (
	ui "conway/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Panel struct {
	*ui.UIBase
	Color rl.Color
}

func NewPanel(color rl.Color) *Panel {
	return &Panel{
		UIBase: ui.NewUIBase(),
		Color:  color,
	}
}

func (p *Panel) GetUIBase() *ui.UIBase {
	return p.UIBase
}

func (p *Panel) Draw() {
	rl.DrawRectangleRec(p.Bounds, p.Color)
}

func (p *Panel) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, p.Bounds)
}
