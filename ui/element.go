package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Element interface {
	GetUIBase() *UIBase
}

type Drawable interface {
	Draw()
}

type Hoverable interface {
	IsHovered(mouse rl.Vector2) bool
}
