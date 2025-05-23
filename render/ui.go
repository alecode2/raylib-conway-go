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
	font := assets.LoadFont("./assets/Font/RobotoMono-Medium.ttf", 96)
	slicetext := assets.LoadTexture("./assets/9slice.png")

	panelStyle := &ui.StyleSheet{
		States: map[ui.UIState]ui.StyleSet{
			ui.UIStateDefault:  {ui.Tint: rl.White},
			ui.UIStateHovered:  {ui.Tint: rl.White},
			ui.UIStatePressed:  {ui.Tint: rl.White},
			ui.UIStateDisabled: {ui.Tint: rl.White},
		},
		Animations: map[ui.StyleProperty]ui.AnimationConfig{
			ui.Tint: {Duration: 0.2, Easing: ui.EaseOutQuad},
		},
	}

	buttonStyle := &ui.StyleSheet{
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

	labelStyle := &ui.StyleSheet{
		States: map[ui.UIState]ui.StyleSet{
			ui.UIStateDefault: {ui.Tint: rl.White},
		},
		Animations: map[ui.StyleProperty]ui.AnimationConfig{
			ui.Tint: {Duration: 0.2, Easing: ui.EaseOutQuad},
		},
	}

	root := cmp.NewContainer()
	root.ID = "ROOT"
	root.Direction = ui.Horizontal
	root.Width = float32(state.ScreenWidth)
	root.Height = float32(state.ScreenHeight)
	root.WidthSizing = ui.SizingFixed
	root.HeightSizing = ui.SizingFixed
	root.MainAlign = ui.AlignCenter
	root.CrossAlign = ui.CrossAlignCenter
	root.Visible = true

	panel := cmp.NewImagePanel(slicetext, panelStyle)
	panel.ID = "PAUSE_PANEL"
	panel.Direction = ui.Vertical
	panel.WidthSizing = ui.SizingFit
	panel.HeightSizing = ui.SizingFit
	panel.Visible = false
	panel.MainAlign = ui.AlignCenter
	panel.CrossAlign = ui.CrossAlignCenter
	panel.PaddingBottom = 16
	panel.PaddingLeft = 16
	panel.PaddingRight = 16
	panel.Gap = 8
	panel.DrawConfig = ui.DrawConfig{
		Mode:       ui.DrawModeNineSlice,
		NineSlice:  ui.MakeNineSliceRegions(slicetext, 16, 80, 16, 80),
		TileCenter: true,
		TileEdges:  true,
	}

	button := cmp.NewImageButton(slicetext, buttonStyle)
	button.ID = "RESUME_BTN"
	button.Width = 192
	button.Height = 64
	button.WidthSizing = ui.SizingFixed
	button.HeightSizing = ui.SizingFixed
	button.MainAlign = ui.AlignCenter
	button.CrossAlign = ui.CrossAlignCenter
	button.Visible = true
	button.DrawConfig = ui.DrawConfig{
		Mode:       ui.DrawModeNineSlice,
		NineSlice:  ui.MakeNineSliceRegions(slicetext, 16, 80, 16, 80),
		TileCenter: true,
		TileEdges:  true,
	}

	ui.AddEventHandler(button, ui.EventRelease, func(ui.UIEvent) {
		bus.Emit(event.Event{Name: "toggle_pause"})
	})

	label := cmp.NewLabel("GAME PAUSED", font, 64, labelStyle)
	label.ID = "PAUSE_LABEL"
	label.WidthSizing = ui.SizingFit
	label.HeightSizing = ui.SizingFit
	label.TextAlign = ui.AlignTextCenter
	label.Wrap = false
	label.Spacing = 1

	btnlabel := cmp.NewLabel("RESUME", font, 48, labelStyle)
	btnlabel.ID = "RESUME_LABEL"
	btnlabel.WidthSizing = ui.SizingFit
	btnlabel.HeightSizing = ui.SizingFit
	btnlabel.TextAlign = ui.AlignTextCenter
	btnlabel.Wrap = false
	btnlabel.Spacing = 0

	ui.AddChild(root, panel)
	ui.AddChild(panel, label)
	ui.AddChild(panel, button)
	ui.AddChild(button, btnlabel)

	uiMap := map[string]ui.Element{
		"ROOT":         root,
		"PAUSE_PANEL":  panel,
		"RESUME_BTN":   button,
		"RESUME_LABEL": btnlabel,
		"PAUSE_LABEL":  label,
	}

	return root, uiMap
}
