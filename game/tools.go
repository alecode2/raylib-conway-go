package game

import (
	"conway/event"
	"fmt"
)

type Tool string

const (
	Paint      Tool = "tool_paint"
	Erase      Tool = "tool_erase"
	Eyedropper Tool = "tool_eyedropper"
)

func InitToolBox(state *GameState, settings *Settings, bus *event.EventBus) {

	bus.Subscribe("toggle_pause", func(e event.Event) {
		state.IsPaused = !state.IsPaused
		fmt.Println("pause toggle called Toolbox function")
	})

	bus.Subscribe("step_forward", func(e event.Event) {
		state.StepForward = true
		fmt.Printf("Step Forward called\n")
	})

	bus.Subscribe("select_tool_paint", func(e event.Event) {
		fmt.Printf("selecting paint tool\n")
		state.ActiveTool = Paint
	})

	bus.Subscribe("select_tool_erase", func(e event.Event) {
		fmt.Printf("selecting erase tool\n")
		state.ActiveTool = Erase
	})

	bus.Subscribe("select_tool_eyedropper", func(e event.Event) {
		fmt.Printf("selecting eyedropper tool\n")
		state.ActiveTool = Eyedropper
	})

	bus.Subscribe("request_map_save", func(e event.Event) {
		fmt.Printf("requesting map save\n")
	})

	bus.Subscribe("request_map_load", func(e event.Event) {
		fmt.Printf("requesting map load\n")
	})

	bus.Subscribe("request_map_export", func(e event.Event) {
		fmt.Printf("requesting map export\n")
	})
}
