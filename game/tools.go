package game

import (
	"conway/event"
	"fmt"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	bus.Subscribe("hex_input_submit", func(e event.Event) {
		fmt.Printf("submitted hex value is :%s\n", e.Data)

		str, ok := e.Data.(string)
		if !ok {
			fmt.Printf("submitted hex is invalid\n")
			return
		}

		color := ColorFromHex(str)
		fmt.Printf("rl.Color value %v\n", color)

		state.SelectedColor = color
	})
}

/*
We convert to RGB and create a rl.Color object with alpha at 255 always
*/
func ColorFromHex(hex string) rl.Color {
	hex = strings.TrimPrefix(hex, "#")

	if len(hex) != 6 {
		return rl.White
	}

	r64, err := strconv.ParseInt(hex[0:2], 16, 0)
	if err != nil {
		return rl.White
	}
	g64, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		return rl.White
	}
	b64, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		return rl.White
	}

	return rl.Color{R: uint8(r64), G: uint8(g64), B: uint8(b64), A: uint8(255)}
}
