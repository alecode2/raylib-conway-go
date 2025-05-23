package cmp

import (
	ui "conway/ui"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Image struct {
	*ui.UIBase
	Texture    rl.Texture2D
	DrawConfig ui.DrawConfig
}

func (i *Image) GetUIBase() *ui.UIBase {
	return i.UIBase
}

func (i *Image) Draw() {
	if i.Texture.ID == 0 {
		fmt.Println("WARNING: No Texture ID")
		return
	}

	tintVal := ui.ResolveStyle(i.UIBase, ui.Tint)
	tint, ok := tintVal.(rl.Color)
	if !ok {
		tint = rl.White
	}

	dest := i.Bounds
	switch i.DrawConfig.Mode {
	case ui.DrawModeSimple:
		src := rl.NewRectangle(0, 0, float32(i.Texture.Width), float32(i.Texture.Height))
		rl.DrawTexturePro(i.Texture, src, dest, rl.NewVector2(0, 0), 0, tint)

	case ui.DrawModeNineSlice:
		ui.DrawNineSlice(
			i.Texture,
			i.DrawConfig.NineSlice,
			dest,
			tint,
			i.DrawConfig.TileCenter,
			i.DrawConfig.TileEdges,
		)

	case ui.DrawModeTiled:
		src := rl.NewRectangle(0, 0, float32(i.Texture.Width), float32(i.Texture.Height))
		ui.TileTexture(i.Texture, src, dest, tint)

	default:
		fmt.Println("ERROR: Unknown DrawMode")
	}
}

func (i *Image) Measure(axis ui.Axis) float32 {
	if axis == ui.Horizontal {
		return float32(i.Texture.Width)
	}
	return float32(i.Texture.Height)
}

func (i *Image) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, i.Bounds)
}
