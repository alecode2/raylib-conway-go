package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Text struct {
	UIBase
	Text    string
	Size    float32
	Color   rl.Color
	Font    rl.Font
	Spacing float32
}

func (t *Text) Draw() {
	if !t.UIBase.Visible {
		return
	}

	pos := rl.NewVector2(t.UIBase.Bounds.X, t.UIBase.Bounds.Y)

	rl.DrawTextEx(t.Font, t.Text, pos, t.Size, t.Spacing, t.Color)

	for _, child := range t.UIBase.Children {
		child.Draw()
	}
}

func (t *Text) IsVisible() bool {
	return t.UIBase.Visible
}

func (t *Text) SetVisible(visible bool) {
	t.UIBase.Visible = visible
}

func (t *Text) AddChild(child Element) {
	child.SetParent(t)
	t.UIBase.AddChild(child)
}

func (t *Text) GetChildren() []Element {
	return t.UIBase.Children
}

func (t *Text) SetParent(parent Element) {
	t.UIBase.SetParent(parent)
}

func (t *Text) GetParent() Element {
	return t.UIBase.GetParent()
}

func (t *Text) HandleEvent(event UIEvent) {
	t.UIBase.HandleEvent(event)
}

func (t *Text) SetEventHandler(eventName string, handler func(UIEvent)) {
	t.UIBase.SetEventHandler(eventName, handler)
}

func (t *Text) HasEventHandler(eventName string) bool {
	return t.UIBase.EventHandlers[eventName] != nil
}

func (t *Text) ShouldPropagate() bool {
	return t.UIBase.PropagateEvents
}

func (t *Text) GetBounds() rl.Rectangle {
	size := rl.MeasureTextEx(t.Font, t.Text, t.Size, t.Spacing)
	return rl.Rectangle{
		X:      t.UIBase.Bounds.X,
		Y:      t.UIBase.Bounds.Y,
		Width:  size.X,
		Height: size.Y,
	}
}

func (t *Text) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, t.GetBounds())
}
