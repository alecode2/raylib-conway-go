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

func InitUI(state GameState, settings Settings, bus *EventBus) (ui.Element, map[string]ui.Element) {

	root := &ui.Container{
		UIBase: &ui.UIBase{
			ID:           "ROOT",
			Direction:    ui.Horizontal,
			Width:        float32(state.ScreenWidth),
			Height:       float32(state.ScreenHeight),
			WidthSizing:  ui.SizingFixed,
			HeightSizing: ui.SizingFixed,
			MainAlign:    ui.AlignCenter,
			CrossAlign:   ui.CrossAlignCenter,
			Visible:      true,
		},
	}

	label := &ui.Label{
		UIBase: &ui.UIBase{
			ID:           "GREETING_LABEL",
			WidthSizing:  ui.SizingFit,
			HeightSizing: ui.SizingFit,
			Visible:      false,
		},
		Text:      "PAUSED",
		Font:      rl.GetFontDefault(),
		FontSize:  64,
		FontColor: rl.Black,
		TextAlign: ui.AlignTextCenter,
		Wrap:      false,
		Spacing:   float32(1),
	}

	ui.AddChild(root, label)

	//Populating the ui registry
	uiMap := make(map[string]ui.Element)
	uiMap["ROOT"] = root
	uiMap["GREETING_LABEL"] = label

	return root, uiMap
}
