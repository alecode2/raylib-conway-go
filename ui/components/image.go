package cmp

import (
	ui "conway/ui"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Image struct {
	*ui.UIBase
	Texture      rl.Texture2D
	TintDefault  rl.Color
	TintHovered  rl.Color
	TintPressed  rl.Color
	TintDisabled rl.Color
}

func NewImage(texture rl.Texture2D, tint rl.Color) *Image {
	return &Image{
		UIBase:       ui.NewUIBase(),
		Texture:      texture,
		TintDefault:  tint,
		TintHovered:  tint,
		TintPressed:  tint,
		TintDisabled: tint,
	}
}

func (i *Image) GetUIBase() *ui.UIBase {
	return i.UIBase
}

func (i *Image) Draw() {
	if i.Texture.ID == 0 {
		fmt.Printf("WARNING: No Texture ID\n")
		return
	}

	var tint rl.Color
	switch i.UIBase.State {
	case ui.UIStateHovered:
		tint = i.TintHovered
	case ui.UIStatePressed:
		tint = i.TintPressed
	case ui.UIStateDisabled:
		tint = i.TintDisabled
	default:
		tint = i.TintDefault
	}

	src := rl.NewRectangle(0, 0, float32(i.Texture.Width), float32(i.Texture.Height))
	dest := ui.GetBounds(i)
	origin := rl.NewVector2(0, 0)
	rotation := float32(0)

	rl.DrawTexturePro(i.Texture, src, dest, origin, rotation, tint)
}

func (i *Image) Measure(axis ui.Axis) float32 {
	if axis == ui.Horizontal {
		return float32(i.Texture.Width)
	}
	return float32(i.Texture.Height)
}
