package render

import (
	"conway/assets"
	"conway/event"
	"conway/game"
	ui "conway/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Event = event.Event
type EventBus = event.EventBus
type GameState = game.GameState
type Settings = game.Settings

func InitUI() ui.Element {

	root := &ui.Panel{
		UIBase: &ui.UIBase{
			ID:            "ROOT",
			Direction:     ui.Horizontal,
			Width:         800,
			Height:        600,
			PaddingTop:    16,
			PaddingBottom: 16,
			PaddingLeft:   16,
			PaddingRight:  16,
			Gap:           8,
			WidthSizing:   ui.SizingFixed,
			HeightSizing:  ui.SizingFixed,
			MainAlign:     ui.AlignCenter,
			CrossAlign:    ui.CrossAlignCenter,
		},
		Color: rl.LightGray,
	}

	// Load the texture from file
	texture := assets.LoadTexture("assets/PNG/Blue/Default/star.png") // Make sure the file exists

	icon := &ui.Image{
		UIBase: &ui.UIBase{
			ID:           "ICON",
			Width:        32,
			Height:       30,
			WidthSizing:  ui.SizingFixed,
			HeightSizing: ui.SizingFixed,
		},
		Texture: texture,
		Tint:    rl.White,
	}
	ui.AddChild(root, icon)

	label := &ui.Label{
		UIBase: &ui.UIBase{
			ID:           "GREETING_LABEL",
			WidthSizing:  ui.SizingFit,
			HeightSizing: ui.SizingFit,
		},
		Text:      "PAUSED",
		Font:      rl.GetFontDefault(),
		FontSize:  24,
		FontColor: rl.Black,
		TextAlign: ui.AlignTextCenter,
		Wrap:      false,
		Spacing:   float32(1),
	}

	ui.AddChild(root, label)

	icon2 := &ui.Image{
		UIBase: &ui.UIBase{
			ID:           "ICON",
			Width:        32,
			Height:       30,
			WidthSizing:  ui.SizingFixed,
			HeightSizing: ui.SizingFixed,
		},
		Texture: texture,
		Tint:    rl.White,
	}
	ui.AddChild(root, icon2)

	return root
}
