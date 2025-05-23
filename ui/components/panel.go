package cmp

import (
	ui "conway/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Panel struct {
	*ui.UIBase
}

func NewPanel(color rl.Color) *Panel {
	return &Panel{
		UIBase: ui.NewUIBase(),
	}
}

func (p *Panel) GetUIBase() *ui.UIBase {
	return p.UIBase
}

func (p *Panel) Draw() {
	colorVal := ui.ResolveStyle(p.UIBase, ui.Tint)
	color, ok := colorVal.(rl.Color)
	if !ok {
		color = rl.White
	}

	rl.DrawRectangleRec(p.Bounds, color)
}

func (p *Panel) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, p.Bounds)
}
