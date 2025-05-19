package cmp

import (
	"conway/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ImagePanel struct {
	*ui.UIBase
	Texture      rl.Texture2D
	TintDefault  rl.Color
	TintHovered  rl.Color
	TintPressed  rl.Color
	TintDisabled rl.Color

	DrawConfig ui.DrawConfig
}

func (ip *ImagePanel) GetUIBase() *ui.UIBase {
	return ip.UIBase
}

func (ip *ImagePanel) Draw() {
	//Check for texture in memory
	if ip.Texture.ID == 0 {
		fmt.Printf("WARNING: No Texture ID\n")
		return
	}

	//Figure out the tint
	var tint rl.Color
	switch ip.GetUIBase().State {
	case ui.UIStateDisabled:
		tint = ip.TintDisabled
	case ui.UIStateHovered:
		tint = ip.TintHovered
	case ui.UIStatePressed:
		tint = ip.TintPressed
	default:
		tint = ip.TintDefault
	}

	switch ip.DrawConfig.Mode {
	case ui.DrawModeSimple:
		src := rl.NewRectangle(0, 0, float32(ip.Texture.Width), float32(ip.Texture.Height))
		rl.DrawTexturePro(ip.Texture, src, ip.Bounds, rl.NewVector2(0, 0), 0, tint)

	case ui.DrawModeNineSlice:
		ui.DrawNineSlice(
			ip.Texture,
			ip.DrawConfig.NineSlice,
			ip.Bounds,
			tint,
			ip.DrawConfig.TileCenter,
			ip.DrawConfig.TileEdges,
		)

	case ui.DrawModeTiled:
		src := rl.NewRectangle(0, 0, float32(ip.Texture.Width), float32(ip.Texture.Height))
		ui.TileTexture(ip.Texture, src, ip.Bounds, tint)

	// You can expand this easily
	default:
		fmt.Println("ERROR: Unknown DrawMode")
	}
}

func (ip *ImagePanel) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, ip.Bounds)
}
