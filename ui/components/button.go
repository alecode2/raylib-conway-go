package cmp

import (
	ui "conway/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	*ui.UIBase
	Color         rl.Color
	ColorPressed  rl.Color
	ColorHovered  rl.Color
	ColorDisabled rl.Color
}

func NewButton(color rl.Color) *Button {
	return &Button{
		UIBase:        ui.NewUIBase(),
		Color:         color,
		ColorPressed:  color,
		ColorHovered:  color,
		ColorDisabled: color,
	}
}

func (b *Button) GetUIBase() *ui.UIBase {
	return b.UIBase
}

func (b *Button) Draw() {
	base := b.UIBase
	var color rl.Color

	switch base.State {
	case ui.UIStateDisabled:
		color = b.ColorDisabled
	case ui.UIStateHovered:
		color = b.ColorHovered
	case ui.UIStatePressed:
		color = b.ColorPressed
	default:
		color = b.Color
	}

	rl.DrawRectangleRec(b.Bounds, color)
}

func (b *Button) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, b.Bounds)
}
