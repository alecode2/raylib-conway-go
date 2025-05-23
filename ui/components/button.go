package cmp

import (
	ui "conway/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	*ui.UIBase
}

func NewButton(color rl.Color) *Button {
	return &Button{
		UIBase: ui.NewUIBase(),
	}
}

func (b *Button) GetUIBase() *ui.UIBase {
	return b.UIBase
}

func (b *Button) Draw() {
	// Resolve tint from style (animated or not)
	tintVal := ui.ResolveStyle(b.UIBase, ui.Tint)
	tint, ok := tintVal.(rl.Color)
	if !ok {
		tint = rl.White
	}

	rl.DrawRectangleRec(b.Bounds, tint)
}

func (b *Button) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, b.Bounds)
}
