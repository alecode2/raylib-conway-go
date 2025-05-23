package cmp

import (
	ui "conway/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ImageButton struct {
	*ui.UIBase
	Texture    rl.Texture2D
	DrawConfig ui.DrawConfig
}

func NewImageButton(texture rl.Texture2D, style *ui.StyleSheet) *ImageButton {
	base := ui.NewUIBase()
	base.Style = style
	return &ImageButton{
		UIBase:     base,
		Texture:    texture,
		DrawConfig: ui.DrawConfig{Mode: ui.DrawModeSimple},
	}
}

func (ib *ImageButton) GetUIBase() *ui.UIBase {
	return ib.UIBase
}

func (ib *ImageButton) Draw() {
	if ib.Texture.ID == 0 {
		fmt.Println("WARNING: ImageButton has no texture")
		return
	}

	// Resolve tint from style (animated or not)
	tintVal := ui.ResolveStyle(ib.UIBase, ui.Tint)
	tint, ok := tintVal.(rl.Color)
	if !ok {
		tint = rl.White
	}

	dest := ib.Bounds
	switch ib.DrawConfig.Mode {
	case ui.DrawModeSimple:
		src := rl.NewRectangle(0, 0, float32(ib.Texture.Width), float32(ib.Texture.Height))
		rl.DrawTexturePro(ib.Texture, src, dest, rl.NewVector2(0, 0), 0, tint)

	case ui.DrawModeNineSlice:
		ui.DrawNineSlice(
			ib.Texture,
			ib.DrawConfig.NineSlice,
			dest,
			tint,
			ib.DrawConfig.TileCenter,
			ib.DrawConfig.TileEdges,
		)

	case ui.DrawModeTiled:
		src := rl.NewRectangle(0, 0, float32(ib.Texture.Width), float32(ib.Texture.Height))
		ui.TileTexture(ib.Texture, src, dest, tint)

	default:
		fmt.Println("ERROR: Unknown DrawMode")
	}
}

func (ib *ImageButton) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, ib.Bounds)
}
