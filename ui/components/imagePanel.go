package cmp

import (
	ui "conway/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ImagePanel struct {
	*ui.UIBase
	Texture    rl.Texture2D
	DrawConfig ui.DrawConfig
}

func NewImagePanel(texture rl.Texture2D, style *ui.StyleSheet) *ImagePanel {
	base := ui.NewUIBase()

	return &ImagePanel{
		UIBase:     base,
		Texture:    texture,
		DrawConfig: ui.DrawConfig{Mode: ui.DrawModeSimple},
	}
}

func (ip *ImagePanel) GetUIBase() *ui.UIBase {
	return ip.UIBase
}

func (ip *ImagePanel) Draw() {
	if ip.Texture.ID == 0 {
		fmt.Println("WARNING: No Texture ID")
		return
	}

	tintVal := ui.ResolveStyle(ip.UIBase, ui.Tint)
	tint, ok := tintVal.(rl.Color)
	if !ok {
		tint = rl.White
	}

	dest := ip.Bounds
	switch ip.DrawConfig.Mode {
	case ui.DrawModeSimple:
		src := rl.NewRectangle(0, 0, float32(ip.Texture.Width), float32(ip.Texture.Height))
		rl.DrawTexturePro(ip.Texture, src, dest, rl.NewVector2(0, 0), 0, tint)

	case ui.DrawModeNineSlice:
		ui.DrawNineSlice(
			ip.Texture,
			ip.DrawConfig.NineSlice,
			dest,
			tint,
			ip.DrawConfig.TileCenter,
			ip.DrawConfig.TileEdges,
		)

	case ui.DrawModeTiled:
		src := rl.NewRectangle(0, 0, float32(ip.Texture.Width), float32(ip.Texture.Height))
		ui.TileTexture(ip.Texture, src, dest, tint)

	default:
		fmt.Println("ERROR: Unknown DrawMode")
	}
}

func (ip *ImagePanel) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, ip.Bounds)
}
