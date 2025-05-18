package cmp

import (
	"conway/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ImageButton struct {
	*ui.UIBase
	Texture      rl.Texture2D
	TintDefault  rl.Color
	TintHovered  rl.Color
	TintPressed  rl.Color
	TintDisabled rl.Color

	DrawConfig ui.DrawConfig
}

func (ib *ImageButton) GetUIBase() *ui.UIBase {
	return ib.UIBase
}

func (ib *ImageButton) Draw() {
	//Check for texture in memory
	if ib.Texture.ID == 0 {
		fmt.Printf("WARNING: No Texture ID\n")
		return
	}

	//Figure out the tint
	var tint rl.Color
	switch ib.GetUIBase().State {
	case ui.UIStateDisabled:
		tint = ib.TintDisabled
	case ui.UIStateHovered:
		tint = ib.TintHovered
	case ui.UIStatePressed:
		tint = ib.TintPressed
	default:
		tint = ib.TintDefault
	}

	switch ib.DrawConfig.Mode {
	case ui.DrawModeSimple:
		src := rl.NewRectangle(0, 0, float32(ib.Texture.Width), float32(ib.Texture.Height))
		rl.DrawTexturePro(ib.Texture, src, ib.Bounds, rl.NewVector2(0, 0), 0, tint)

	case ui.DrawModeNineSlice:
		ui.DrawNineSlice(
			ib.Texture,
			ib.DrawConfig.NineSlice,
			ib.Bounds,
			tint,
			ib.DrawConfig.TileCenter,
			ib.DrawConfig.TileEdges,
		)

	case ui.DrawModeTiled:
		src := rl.NewRectangle(0, 0, float32(ib.Texture.Width), float32(ib.Texture.Height))
		ui.TileTexture(ib.Texture, src, ib.Bounds, tint)

	// You can expand this easily
	default:
		fmt.Println("ERROR: Unknown DrawMode")
	}
}

func (ib *ImageButton) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, ib.Bounds)
}
