package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Event types
const (
	EventHover   = "hover"
	EventClick   = "click"
	EventFocus   = "focus"
	EventBlur    = "blur"
	EventPress   = "press"
	EventRelease = "release"
)

type UIEvent struct {
	Name          string
	MousePosition rl.Vector2
}

type UIEventElement struct {
	Element Element
	Depth   int // Used for reverse depth-based handling
}

var uiEventList []UIEventElement

var currentPressedElement Element

func RefreshUIEventList(root Element) {
	uiEventList = uiEventList[:0]

	var dfs func(el Element, depth int)
	dfs = func(el Element, depth int) {
		base := el.GetUIBase()
		if !base.Visible || base.State == UIStateDisabled {
			return
		}

		uiEventList = append(uiEventList, UIEventElement{Element: el, Depth: depth})

		for _, child := range base.Children {
			dfs(child, depth+1)
		}
	}

	dfs(root, 0)
}

func HandleUIHover(mouse rl.Vector2) {
	for i := len(uiEventList) - 1; i >= 0; i-- {
		elem := uiEventList[i].Element
		hoverable, ok := elem.(Hoverable)
		if !ok {
			continue
		}

		hovered := hoverable.IsHovered(mouse)
		base := elem.GetUIBase()

		if hovered {
			if base.State != UIStateHovered && base.State != UIStatePressed {
				SetState(elem, UIStateHovered)
				dispatchEvent(elem, UIEvent{Name: EventHover, MousePosition: mouse})
			}
		} else if base.State == UIStateHovered {
			SetState(elem, UIStateDefault)
			dispatchEvent(elem, UIEvent{Name: EventBlur, MousePosition: mouse})
		}
	}
}

func HandleUIPress(mouse rl.Vector2) bool {
	for i := len(uiEventList) - 1; i >= 0; i-- {
		elem := uiEventList[i].Element
		if hoverable, ok := elem.(Hoverable); ok && hoverable.IsHovered(mouse) {
			currentPressedElement = elem
			dispatchPressFrom(elem)
			return true
		}
	}
	return false
}

func HandleUIClick(mouse rl.Vector2) bool {
	for i := len(uiEventList) - 1; i >= 0; i-- {
		elem := uiEventList[i].Element
		if hoverable, ok := elem.(Hoverable); ok && hoverable.IsHovered(mouse) {
			dispatchClickFrom(elem)
			return true
		}
	}
	return false
}

func HandleUIRelease(mouse rl.Vector2) bool {
	if currentPressedElement != nil {
		dispatchReleaseFrom(currentPressedElement)

		if hoverable, ok := currentPressedElement.(Hoverable); ok && hoverable.IsHovered(mouse) {
			dispatchClickFrom(currentPressedElement)
		}

		currentPressedElement = nil
		return true
	}
	return false
}

func dispatchPressFrom(elem Element) {
	dispatchEvent(elem, UIEvent{
		Name:          EventPress,
		MousePosition: rl.GetMousePosition(),
	})

	SetState(elem, UIStatePressed)

	if elem.GetUIBase().PropagateEvents && elem.GetUIBase().Parent != nil {
		dispatchPressFrom(elem.GetUIBase().Parent)
	}
}

func dispatchClickFrom(elem Element) {
	dispatchEvent(elem, UIEvent{
		Name:          EventClick,
		MousePosition: rl.GetMousePosition(),
	})

	if elem.GetUIBase().PropagateEvents && elem.GetUIBase().Parent != nil {
		dispatchClickFrom(elem.GetUIBase().Parent)
	}
}

func dispatchReleaseFrom(elem Element) {
	dispatchEvent(elem, UIEvent{
		Name:          EventRelease,
		MousePosition: rl.GetMousePosition(),
	})

	SetState(elem, UIStateDefault)

	if elem.GetUIBase().PropagateEvents && elem.GetUIBase().Parent != nil {
		dispatchReleaseFrom(elem.GetUIBase().Parent)
	}
}

func dispatchEvent(elem Element, event UIEvent) {
	base := elem.GetUIBase()
	if handler, ok := base.EventHandlers[event.Name]; ok {
		handler(event)
	}
}
