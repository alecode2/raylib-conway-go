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
}

func (ib *ImageButton) GetUIBase() *ui.UIBase {
	return ib.UIBase
}

func (ib *ImageButton) Draw() {
	if ib.Texture.ID == 0 {
		fmt.Printf("WARNING: No Texture ID\n")
		return
	}

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

	src := rl.NewRectangle(0, 0, float32(ib.Texture.Width), float32(ib.Texture.Height))
	dest := ui.GetBounds(ib)
	origin := rl.NewVector2(0, 0)
	rotation := float32(0)

	rl.DrawTexturePro(ib.Texture, src, dest, origin, rotation, tint)
}

func (ib *ImageButton) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, ib.Bounds)
}
