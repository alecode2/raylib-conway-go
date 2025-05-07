package ui

import rl "github.com/gen2brain/raylib-go/raylib"

// UI TYPES
type UIState string

const (
	UIStateDefault  UIState = "default"
	UIStateHovered  UIState = "hovered"
	UIStatePressed  UIState = "pressed"
	UIStateDisabled UIState = "disabled"
)

/*
UI BASE TYPE
*/
type UIBase struct {
	ID              string
	Visible         bool
	Children        []Element
	Parent          Element
	State           UIState
	EventHandlers   map[string]func(UIEvent)
	PropagateEvents bool
	Bounds          rl.Rectangle
}

func (b *UIBase) SetParent(parent Element) {
	b.Parent = parent
}

func (b *UIBase) GetParent() Element {
	return b.Parent
}

func (b *UIBase) AddChild(child Element) {
	b.Children = append(b.Children, child)
}

func (b *UIBase) SetID(id string) {
	b.ID = id
}

func (b *UIBase) GetID() string {
	return b.ID
}

func (b *UIBase) GetState() UIState {
	return b.State
}

func (b *UIBase) SetState(state UIState) {
	b.State = state
}

func (b *UIBase) SetEventHandler(eventName string, handler func(UIEvent)) {
	if b.EventHandlers == nil {
		b.EventHandlers = make(map[string]func(UIEvent))
	}

	b.EventHandlers[eventName] = handler
}

func (b *UIBase) HandleEvent(event UIEvent) {
	if handler, ok := b.EventHandlers[event.Name]; ok {
		handler(event)
	}
}
