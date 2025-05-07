package render

import (
	"conway/event"
	"conway/game"
	ui "conway/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Event = event.Event
type EventBus = event.EventBus
type GameState = game.GameState
type Settings = game.Settings

func GetUI(state game.GameState, settings Settings, bus *EventBus) (ui.Element, map[string]ui.Element) {
	uiRegistry := make(map[string]ui.Element)

	// Pause Panel
	pausePanel := &ui.Panel{
		Color: rl.Color{R: 255, G: 255, B: 255, A: 127},
		UIBase: ui.UIBase{
			Visible: false,
			Bounds: rl.Rectangle{
				X:      float32(state.ScreenWidth/2 - 256/2),
				Y:      float32(state.ScreenHeight/2 - 128/2),
				Width:  256,
				Height: 192,
			},
		},
	}
	pausePanel.SetID("pause_panel")
	uiRegistry["pause_panel"] = pausePanel

	// Pause Text
	pauseText := &ui.Text{
		Text:    "Paused",
		Size:    64,
		Color:   rl.Color{R: 0, G: 0, B: 0, A: 127},
		Font:    rl.GetFontDefault(),
		Spacing: 4,
		UIBase:  ui.UIBase{Visible: true},
	}

	measured := pauseText.GetBounds()
	pauseText.UIBase.Bounds = rl.Rectangle{
		X:      pausePanel.Bounds.X + (pausePanel.Bounds.Width / 2) - (measured.X / 2),
		Y:      pausePanel.Bounds.Y + (pausePanel.Bounds.Height / 2) - (measured.Y / 2),
		Width:  measured.Width,
		Height: measured.Height,
	}

	pauseText.SetID("pause_text")
	uiRegistry["pause_text"] = pauseText
	pausePanel.AddChild(pauseText)

	// Pause Button
	pauseButton := &ui.Button{
		BgColor: rl.DarkGray,
		UIBase: ui.UIBase{Visible: true,
			Bounds: rl.Rectangle{
				X:      float32(state.ScreenWidth/2 - 80),
				Y:      float32(state.ScreenHeight/2 + 64),
				Width:  160,
				Height: 40,
			},
		},
	}

	pauseButton.SetID("pause_button")
	uiRegistry["pause_button"] = pauseButton
	pausePanel.AddChild(pauseButton)

	// Add OnClick event to button
	pauseButton.SetEventHandler(ui.EventClick, func(event ui.UIEvent) {
		bus.Emit(Event{Name: "toggle_pause"})
	})

	// Pause Button Text
	pauseBtnText := &ui.Text{
		Text:    "Resume",
		Size:    20,
		Color:   rl.Color{R: 0, G: 0, B: 0, A: 255},
		Font:    rl.GetFontDefault(),
		Spacing: 2,
		UIBase: ui.UIBase{
			Visible:         true,
			PropagateEvents: true,
		},
	}

	textSize := pauseBtnText.GetBounds()
	pauseBtnText.UIBase.Bounds = rl.Rectangle{
		X:      pauseButton.Bounds.X + (pauseButton.Bounds.Width / 2) - (textSize.X / 2),
		Y:      pauseButton.Bounds.Y + (pauseButton.Bounds.Height / 2) - (textSize.Y / 2),
		Width:  textSize.Width,
		Height: textSize.Height,
	}

	pauseBtnText.SetID("pause_button_text")
	uiRegistry["pause_button_text"] = pauseBtnText
	pauseButton.AddChild(pauseBtnText)

	return pausePanel, uiRegistry
}

func DrawUI(element ui.Element) {
	if !element.IsVisible() {
		return
	}

	element.Draw()

	for _, child := range element.GetChildren() {
		DrawUI(child)
	}
}
