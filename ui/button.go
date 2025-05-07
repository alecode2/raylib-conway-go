package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	UIBase
	BgColor rl.Color
}

func (b *Button) Draw() {
	if !b.UIBase.Visible {
		return
	}

	bg := b.BgColor
	switch b.GetState() {
	case UIStateHovered:
		bg = rl.Fade(b.BgColor, 0.8)
	case UIStatePressed:
		bg = rl.Red
	case UIStateDisabled:
		bg = rl.Gray
	}

	rl.DrawRectangleRec(b.Bounds, bg)

	for _, child := range b.Children {
		child.Draw()
	}
}

func (b *Button) SetContent(child Element) {
	child.SetParent(b)
	b.Children = []Element{child}
}

func (b *Button) SetParent(parent Element) {
	b.UIBase.SetParent(parent)
}

func (b *Button) GetParent() Element {
	return b.UIBase.GetParent()
}

func (b *Button) AddChild(child Element) {
	child.SetParent(b)
	b.UIBase.AddChild(child)
}

func (b *Button) GetChildren() []Element {
	return b.UIBase.Children
}

func (b *Button) SetVisible(v bool) {
	b.UIBase.Visible = v
}

func (b *Button) IsVisible() bool {
	return b.UIBase.Visible
}

func (b *Button) SetID(id string) {
	b.UIBase.SetID(id)
}

func (b *Button) GetID() string {
	return b.UIBase.GetID()
}

func (b *Button) HandleEvent(event UIEvent) {
	b.UIBase.HandleEvent(event)
}

func (b *Button) SetEventHandler(eventName string, handler func(UIEvent)) {
	b.UIBase.SetEventHandler(eventName, handler)
}

func (b *Button) HasEventHandler(eventName string) bool {
	return b.UIBase.EventHandlers[eventName] != nil
}

func (b *Button) ShouldPropagate() bool {
	return b.UIBase.PropagateEvents
}

//Bounds Checking
func (b *Button) GetBounds() rl.Rectangle {
	return b.UIBase.Bounds
}

func (b *Button) IsHovered(mouse rl.Vector2) bool {
	bounds := b.GetBounds()
	return rl.CheckCollisionPointRec(mouse, bounds)
}
