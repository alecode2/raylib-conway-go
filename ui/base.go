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
	State           UIState
	Visible         bool
	Bounds          rl.Rectangle
	Parent          Element
	Children        []Element
	EventHandlers   map[string]func(UIEvent)
	PropagateEvents bool
	CustomProps     map[string]interface{}
}

func NewUIBase() UIBase {
	return UIBase{
		State:         UIStateDefault,
		Visible:       true,
		Children:      []Element{},
		EventHandlers: make(map[string]func(UIEvent)),
		CustomProps:   make(map[string]interface{}),
	}
}

/*
 UI BASE HELPER FUNCTIONS
*/
func GetState(e Element) UIState {
	return e.GetUIBase().State
}

func SetState(e Element, state UIState) {
	e.GetUIBase().State = state
}

func SetVisible(e Element, visible bool) {
	e.GetUIBase().Visible = visible
}

func IsVisible(e Element) bool {
	return e.GetUIBase().Visible
}

func AddChild(parent, child Element) {
	child.GetUIBase().Parent = parent
	parent.GetUIBase().Children = append(parent.GetUIBase().Children, child)
}

func GetChildren(e Element) []Element {
	return e.GetUIBase().Children
}

func SetID(e Element, id string) {
	e.GetUIBase().ID = id
}

func GetID(e Element) string {
	return e.GetUIBase().ID
}

func SetBounds(e Element, bounds rl.Rectangle) {
	e.GetUIBase().Bounds = bounds
}

func GetBounds(e Element) rl.Rectangle {
	return e.GetUIBase().Bounds
}
