package game

import (
	"conway/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var buttonStyle = &ui.StyleSheet{
	States: map[ui.UIState]ui.StyleSet{
		ui.UIStateDefault:  {ui.Tint: rl.White},
		ui.UIStateHovered:  {ui.Tint: rl.LightGray},
		ui.UIStatePressed:  {ui.Tint: rl.Gray},
		ui.UIStateDisabled: {ui.Tint: rl.Fade(rl.White, 0.5)},
	},
	Animations: map[ui.StyleProperty]ui.AnimationConfig{
		ui.Tint: {Duration: 0.2, Easing: ui.EaseOutQuad},
	},
}

var selectedTool = &ui.StyleSheet{
	States: map[ui.UIState]ui.StyleSet{
		ui.UIStateDefault:  {ui.Tint: rl.Gray},
		ui.UIStateHovered:  {ui.Tint: rl.Gray},
		ui.UIStatePressed:  {ui.Tint: rl.DarkGray},
		ui.UIStateDisabled: {ui.Tint: rl.Fade(rl.White, 0.5)},
	},
	Animations: map[ui.StyleProperty]ui.AnimationConfig{
		ui.Tint: {Duration: 0.2, Easing: ui.EaseOutQuad},
	},
}

var labelStyle = &ui.StyleSheet{
	States: map[ui.UIState]ui.StyleSet{
		ui.UIStateDefault: {ui.Tint: rl.White},
	},
}

var selectedInputField = &ui.StyleSheet{
	/*
		States: map[ui.UIState]ui.StyleSet{
			ui.UIStateDefault:  {ui.Tint: rl.Gray},
			ui.UIStateHovered:  {ui.Tint: rl.Gray},
			ui.UIStatePressed:  {ui.Tint: rl.DarkGray},
			ui.UIStateDisabled: {ui.Tint: rl.Fade(rl.White, 0.5)},
		},
		Animations: map[ui.StyleProperty]ui.AnimationConfig{
			ui.Tint: {Duration: 0.2, Easing: ui.EaseOutQuad},
		},
	*/
}
