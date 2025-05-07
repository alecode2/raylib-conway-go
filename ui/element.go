package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Element interface {
	Draw()
	IsVisible() bool
	SetVisible(bool)
	AddChild(child Element)
	GetChildren() []Element
	SetParent(parent Element)
	GetParent() Element
	SetID(id string)
	GetID() string
	GetState() UIState
	SetState(UIState)
	GetBounds() rl.Rectangle
	IsHovered(mouse rl.Vector2) bool
	//Event Code
	HandleEvent(event UIEvent)
	SetEventHandler(eventName string, handler func(event UIEvent))
	HasEventHandler(eventName string) bool
	ShouldPropagate() bool
}
