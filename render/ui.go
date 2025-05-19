package render

import (
	"conway/assets"
	"conway/event"
	"conway/game"
	ui "conway/ui"
	cmp "conway/ui/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Event = event.Event
type EventBus = event.EventBus
type GameState = game.GameState
type Settings = game.Settings

func InitUI(state GameState, settings Settings, bus *EventBus) (ui.Element, map[string]ui.Element) {
	// Instance the components
	root := &cmp.Container{
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

	slicetext := assets.LoadTexture("./assets/9slice.png")

	panel := &cmp.ImagePanel{
		UIBase: &ui.UIBase{
			ID:            "PAUSE_PANEL",
			Width:         192,
			Height:        64,
			Direction:     ui.Vertical,
			WidthSizing:   ui.SizingFit,
			HeightSizing:  ui.SizingFit,
			EventHandlers: make(map[string]func(ui.UIEvent)),
			MainAlign:     ui.AlignCenter,
			CrossAlign:    ui.CrossAlignCenter,
			Visible:       false,
			PaddingTop:    32,
			PaddingBottom: 32,
		},
		Texture:      slicetext,
		TintDefault:  rl.White,
		TintHovered:  rl.White,
		TintPressed:  rl.White,
		TintDisabled: rl.White,
	}

	panel.DrawConfig = ui.DrawConfig{
		Mode:       ui.DrawModeNineSlice,
		NineSlice:  ui.MakeNineSliceRegions(slicetext, 32, 96, 32, 96),
		TileCenter: true,
		TileEdges:  true,
	}

	//btnTexture := assets.LoadTexture("./assets/PNG/Blue/Default/button_rectangle_depth_gloss.png")

	button := &cmp.ImageButton{
		UIBase: &ui.UIBase{
			ID:            "RESUME_BTN",
			Width:         192,
			Height:        64,
			WidthSizing:   ui.SizingFixed,
			HeightSizing:  ui.SizingFixed,
			EventHandlers: make(map[string]func(ui.UIEvent)),
			MainAlign:     ui.AlignCenter,
			CrossAlign:    ui.CrossAlignCenter,
			Visible:       true,
		},
		Texture:      slicetext,
		TintDefault:  rl.White,
		TintHovered:  rl.LightGray,
		TintPressed:  rl.Gray,
		TintDisabled: rl.Fade(rl.White, 0.5),
	}

	button.DrawConfig = ui.DrawConfig{
		Mode:       ui.DrawModeNineSlice,
		NineSlice:  ui.MakeNineSliceRegions(slicetext, 32, 96, 32, 96),
		TileCenter: true,
		TileEdges:  true,
	}

	ui.AddEventHandler(button, ui.EventRelease, func(ui.UIEvent) {
		bus.Emit(event.Event{Name: "toggle_pause"})
	})

	font := assets.LoadFont("./assets/Font/Kenney Future.ttf")

	label := &cmp.Label{
		UIBase: &ui.UIBase{
			ID:           "PAUSE_LABEL",
			WidthSizing:  ui.SizingFit,
			HeightSizing: ui.SizingFit,
			Visible:      true,
		},
		Text:      "GAME PAUSED",
		Font:      font,
		FontSize:  48,
		FontColor: rl.Black,
		TextAlign: ui.AlignTextCenter,
		Wrap:      false,
		Spacing:   float32(1),
	}

	btnlabel := &cmp.Label{
		UIBase: &ui.UIBase{
			ID:           "RESUME_LABEL",
			WidthSizing:  ui.SizingFit,
			HeightSizing: ui.SizingFit,
			Visible:      true,
		},
		Text:      "RESUME",
		Font:      font,
		FontSize:  32,
		FontColor: rl.White,
		TextAlign: ui.AlignTextCenter,
		Wrap:      false,
		Spacing:   float32(0),
	}

	//Set the tree
	ui.AddChild(root, panel)
	ui.AddChild(panel, label)
	ui.AddChild(panel, button)
	ui.AddChild(button, btnlabel)

	//Populating the ui registry
	uiMap := make(map[string]ui.Element)
	uiMap["ROOT"] = root
	uiMap["PAUSE_PANEL"] = panel
	uiMap["RESUME_BTN"] = button
	uiMap["RESUME_LABEL"] = label

	return root, uiMap
}
