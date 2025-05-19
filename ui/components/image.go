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

	DrawConfig ui.DrawConfig
}

func (i *Image) GetUIBase() *ui.UIBase {
	return i.UIBase
}

func (i *Image) Draw() {
	//Check for texture in memory
	if i.Texture.ID == 0 {
		fmt.Printf("WARNING: No Texture ID\n")
		return
	}

	//Figure out the tint
	var tint rl.Color
	switch i.GetUIBase().State {
	case ui.UIStateDisabled:
		tint = i.TintDisabled
	case ui.UIStateHovered:
		tint = i.TintHovered
	case ui.UIStatePressed:
		tint = i.TintPressed
	default:
		tint = i.TintDefault
	}

	switch i.DrawConfig.Mode {
	case ui.DrawModeSimple:
		src := rl.NewRectangle(0, 0, float32(i.Texture.Width), float32(i.Texture.Height))
		rl.DrawTexturePro(i.Texture, src, i.Bounds, rl.NewVector2(0, 0), 0, tint)

	case ui.DrawModeNineSlice:
		ui.DrawNineSlice(
			i.Texture,
			i.DrawConfig.NineSlice,
			i.Bounds,
			tint,
			i.DrawConfig.TileCenter,
			i.DrawConfig.TileEdges,
		)

	case ui.DrawModeTiled:
		src := rl.NewRectangle(0, 0, float32(i.Texture.Width), float32(i.Texture.Height))
		ui.TileTexture(i.Texture, src, i.Bounds, tint)

	// You can expand this easily
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
