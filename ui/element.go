package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Element interface {
	GetUIBase() *UIBase
}

type Measurable interface {
	// Measure returns the desired size along the given axis,
	// *excluding* padding/gaps.
	Measure(axis Axis) float32
}
type Drawable interface {
	Draw()
}

type Hoverable interface {
	IsHovered(mouse rl.Vector2) bool
}
