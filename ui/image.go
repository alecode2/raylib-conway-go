package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Image struct {
	*UIBase
	Texture rl.Texture2D
	Tint    rl.Color
}

func NewImage(texture rl.Texture2D, tint rl.Color) *Image {
	return &Image{
		UIBase:  NewUIBase(),
		Texture: texture,
		Tint:    tint,
	}
}

func (i *Image) GetUIBase() *UIBase {
	return i.UIBase
}

func (i *Image) Draw() {
	if i.Texture.ID == 0 {
		fmt.Printf("WARNING: No Texture ID\n")
		return
	}

	src := rl.NewRectangle(0, 0, float32(i.Texture.Width), float32(i.Texture.Height))
	dest := GetBounds(i)
	origin := rl.NewVector2(0, 0)
	rotation := float32(0)

	rl.DrawTexturePro(i.Texture, src, dest, origin, rotation, i.Tint)
}

func (i *Image) Measure(axis Axis) float32 {
	if axis == Horizontal {
		return float32(i.Texture.Width)
	}
	return float32(i.Texture.Height)
}
