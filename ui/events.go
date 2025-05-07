package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Types of Events supported by default
const (
	EventHover = "hover"
	EventClick = "click"
	EventFocus = "focus"
	EventBlur  = "blur"
	EventPress = "press"
)

type UIEvent struct {
	Name          string
	MousePosition rl.Vector2
}

type UIEventElement struct {
	Element Element
	Depth   int // For sorting later
}

var uiEventList []UIEventElement

func RefreshUIEventList(root Element) {
	uiEventList = uiEventList[:0] // Clear the existing list without reallocating

	var dfs func(el Element, depth int)
	dfs = func(el Element, depth int) {
		if !el.IsVisible() {
			return
		}

		if el.GetState() != UIStateDisabled {
			uiEventList = append(uiEventList, UIEventElement{Element: el, Depth: depth})
		}

		for _, child := range el.GetChildren() {
			dfs(child, depth+1)
		}
	}

	dfs(root, 0)
}

func HandleUIHover(mouse rl.Vector2) {
	hoveredSomething := false

	for i := len(uiEventList) - 1; i >= 0; i-- {
		elem := uiEventList[i].Element
		if elem.IsHovered(mouse) {
			elem.SetState(UIStateHovered)
			hoveredSomething = true
		} else {
			elem.SetState(UIStateDefault)
		}
	}

	if !hoveredSomething {
	}

}

func HandleUIPress(mouse rl.Vector2) bool {
	for i := len(uiEventList) - 1; i >= 0; i-- {
		elem := uiEventList[i].Element
		if elem.IsHovered(mouse) {
			dispatchPressFrom(elem)
			return true
		}
	}

	return false
}

func dispatchPressFrom(elem Element) {
	elem.HandleEvent(UIEvent{
		Name:          "press",
		MousePosition: rl.GetMousePosition(),
	})

	elem.SetState(UIStatePressed)

	if elem.ShouldPropagate() && elem.GetParent() != nil {
		dispatchPressFrom(elem.GetParent())
	}
}

func HandleUIClick(mouse rl.Vector2) bool {
	for i := len(uiEventList) - 1; i >= 0; i-- {
		elem := uiEventList[i].Element
		if elem.IsHovered(mouse) {
			dispatchClickFrom(elem)
			return true
		}
	}
	return false
}

func dispatchClickFrom(elem Element) {
	elem.HandleEvent(UIEvent{
		Name:          "click",
		MousePosition: rl.GetMousePosition(),
	})

	if elem.ShouldPropagate() && elem.GetParent() != nil {
		dispatchClickFrom(elem.GetParent())
	}
}
