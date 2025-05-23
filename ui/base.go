package ui

import rl "github.com/gen2brain/raylib-go/raylib"

// UI STATES
type UIState string

const (
	UIStateDefault  UIState = "default"
	UIStateHovered  UIState = "hovered"
	UIStatePressed  UIState = "pressed"
	UIStateDisabled UIState = "disabled"
)

//SIZING
type SizingMode int

const (
	SizingFixed SizingMode = iota
	SizingFit
	SizingGrow
)

type Axis int

const (
	Horizontal Axis = iota
	Vertical
)

//POSITIONING
type MainAxisAlignment int

const (
	AlignStart MainAxisAlignment = iota
	AlignCenter
	AlignEnd
	AlignSpaceBetween
	AlignSpaceAround
)

type CrossAxisAlignment int

const (
	CrossAlignStart CrossAxisAlignment = iota
	CrossAlignCenter
	CrossAlignEnd
)

type TextAlign int

const (
	AlignTextLeft = iota
	AlignTextCenter
	AlignTextRight
)

//For texture tiling of ui elements
type DrawMode int

const (
	DrawModeSimple DrawMode = iota
	DrawModeNineSlice
	DrawModeTiled
)

type DrawConfig struct {
	Mode       DrawMode
	NineSlice  [9]rl.Rectangle // Used only if Mode == DrawModeNineSlice
	TileCenter bool            // Only for 9-slice/tile modes
	TileEdges  bool
}

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

	// Layout properties
	Width        float32
	Height       float32
	MinWidth     float32
	MinHeight    float32
	WidthSizing  SizingMode
	HeightSizing SizingMode
	Direction    Axis
	MainAlign    MainAxisAlignment
	CrossAlign   CrossAxisAlignment
	ZIndex       int

	PaddingTop    float32
	PaddingRight  float32
	PaddingBottom float32
	PaddingLeft   float32
	Gap           float32

	//Style things
	Style     *StyleSheet
	AnimState map[StyleProperty]*StyleAnimation
}

func NewUIBase() *UIBase {
	return &UIBase{
		State:         UIStateDefault,
		Visible:       true,
		Children:      []Element{},
		EventHandlers: make(map[string]func(UIEvent)),
		CustomProps:   make(map[string]interface{}),
		WidthSizing:   SizingFixed,
		HeightSizing:  SizingFixed,
		Direction:     Vertical,
	}
}

/*
 UI BASE HELPER FUNCTIONS
*/
func GetState(e Element) UIState {
	return e.GetUIBase().State
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

func AddEventHandler(e Element, eventName string, handler func(UIEvent)) {
	base := e.GetUIBase()
	base.EventHandlers[eventName] = handler
}

func RemoveEventHandler(e Element, eventName string) {
	base := e.GetUIBase()
	delete(base.EventHandlers, eventName)
}
