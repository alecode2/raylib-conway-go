package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Panel struct {
	UIBase
	Color rl.Color
}

func (p *Panel) Draw() {
	if p.UIBase.Visible {
		rl.DrawRectangleRec(p.Bounds, p.Color)
	}
}

func (p *Panel) IsVisible() bool {
	return p.UIBase.Visible
}

func (p *Panel) SetVisible(visible bool) {
	p.UIBase.Visible = visible
}

func (p *Panel) AddChild(child Element) {
	child.SetParent(p)
	p.UIBase.AddChild(child)
}

func (p *Panel) GetChildren() []Element {
	return p.UIBase.Children
}

func (p *Panel) SetParent(parent Element) {
	p.UIBase.SetParent(parent)
}

func (p *Panel) GetParent() Element {
	return p.UIBase.GetParent()
}

func (p *Panel) HandleEvent(event UIEvent) {
	p.UIBase.HandleEvent(event)
}

func (p *Panel) SetEventHandler(eventName string, handler func(UIEvent)) {
	p.UIBase.SetEventHandler(eventName, handler)
}

func (p *Panel) HasEventHandler(eventName string) bool {
	return p.UIBase.EventHandlers[eventName] != nil
}

func (p *Panel) ShouldPropagate() bool {
	return p.UIBase.PropagateEvents
}

func (p *Panel) GetBounds() rl.Rectangle {
	return p.UIBase.Bounds
}

func (p *Panel) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, p.GetBounds())
}
