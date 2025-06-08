package game

import (
	"conway/assets"
	"conway/event"
	ui "conway/ui"
	cmp "conway/ui/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Event = event.Event
type EventBus = event.EventBus

func InitUI(state *GameState, settings *Settings, bus *EventBus) (ui.Element, map[string]ui.Element) {
	uiMap := make(map[string]ui.Element)

	root := root("ROOT", float32(state.ScreenWidth), float32(state.ScreenHeight), uiMap)
	panel := panel("TOOL_PANEL", uiMap)
	header := sectionHeader("TOOL_HEADER", "TOOLS", uiMap)
	resumeBtn := labelButton("RESUME_BUTTON", "RESUME", "toggle_pause", bus, uiMap)
	paintBtn := labelButton("PAINT_BUTTON", "PAINT CELLS", "select_tool_paint", bus, uiMap)
	eraseBtn := labelButton("ERASE_BUTTON", "ERASE CELLS", "select_tool_erase", bus, uiMap)
	dropperBtn := labelButton("EYEDROPPER_BUTTON", "SAMPLE COLOR", "select_tool_eyedropper", bus, uiMap)
	saveBtn := labelButton("SAVE_BUTTON", "SAVE MAP", "request_map_save", bus, uiMap)
	loadBtn := labelButton("LOAD_BUTTON", "LOAD MAP", "request_map_load", bus, uiMap)
	exportBtn := labelButton("EXPORT_BUTTON", "EXPORT MAP", "request_map_export", bus, uiMap)

	ui.AddChild(root, panel)
	ui.AddChild(panel, header)
	ui.AddChild(panel, resumeBtn)
	ui.AddChild(panel, paintBtn)
	ui.AddChild(panel, eraseBtn)
	ui.AddChild(panel, dropperBtn)
	ui.AddChild(panel, saveBtn)
	ui.AddChild(panel, loadBtn)
	ui.AddChild(panel, exportBtn)

	//Here we only subscribe to events in order to do ui changes
	bus.Subscribe("toggle_pause", func(e event.Event) {
		if pauseBtn, ok := uiMap["TOOL_PANEL"]; ok {
			pauseBtn.GetUIBase().Visible = state.IsPaused
		}
	})

	bus.Subscribe("select_tool_paint", func(e event.Event) {
		paintBtn.Style = selectedTool
		eraseBtn.Style = buttonStyle
		dropperBtn.Style = buttonStyle
	})

	bus.Subscribe("select_tool_erase", func(e event.Event) {
		paintBtn.Style = buttonStyle
		eraseBtn.Style = selectedTool
		dropperBtn.Style = buttonStyle
	})

	bus.Subscribe("select_tool_eyedropper", func(e event.Event) {
		paintBtn.Style = buttonStyle
		eraseBtn.Style = buttonStyle
		dropperBtn.Style = selectedTool
	})

	return root, uiMap
}

func root(id string, width, height float32, uiMap map[string]ui.Element) *cmp.Container {
	root := cmp.NewContainer()
	root.ID = id
	root.Direction = ui.Horizontal
	root.Width = float32(width)
	root.Height = float32(height)
	root.WidthSizing = ui.SizingFixed
	root.HeightSizing = ui.SizingFixed
	root.MainAlign = ui.AlignEnd
	root.CrossAlign = ui.CrossAlignCenter
	root.Visible = true
	uiMap["ROOT"] = root
	return root
}

func panel(id string, uiMap map[string]ui.Element) *cmp.ImagePanel {
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

	panel := cmp.NewImagePanel(slicetext, panelStyle)
	panel.ID = id
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
	uiMap[id] = panel
	return panel
}

func sectionHeader(id, content string, uiMap map[string]ui.Element) *cmp.Label {
	font := assets.LoadFont("./assets/Font/RobotoMono-Medium.ttf", 96)

	labelStyle := &ui.StyleSheet{
		States: map[ui.UIState]ui.StyleSet{
			ui.UIStateDefault: {ui.Tint: rl.White},
		},
	}

	label := cmp.NewLabel(content, font, 18, labelStyle)
	label.ID = "TOOL_LABEL"
	label.WidthSizing = ui.SizingFit
	label.HeightSizing = ui.SizingFit
	label.TextAlign = ui.AlignTextCenter
	label.Wrap = false
	label.Spacing = 1

	uiMap[id] = label
	return label
}

func labelButton(id, label, eventName string, bus *EventBus, uiMap map[string]ui.Element) *cmp.ImageButton {
	font := assets.LoadFont("./assets/Font/RobotoMono-Medium.ttf", 96)
	slicetext := assets.LoadTexture("./assets/btn.png")

	button := cmp.NewImageButton(slicetext, buttonStyle)
	button.ID = id
	button.Width = 196
	button.Height = 32
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
		bus.Emit(event.Event{Name: eventName})
	})

	labelID := id + "_LABEL"
	btnlabel := cmp.NewLabel(label, font, 18, labelStyle)
	btnlabel.ID = labelID
	btnlabel.WidthSizing = ui.SizingFit
	btnlabel.HeightSizing = ui.SizingFit
	btnlabel.TextAlign = ui.AlignTextCenter
	btnlabel.Wrap = false
	btnlabel.Spacing = 0

	ui.AddChild(button, btnlabel)
	uiMap[id] = button
	uiMap[labelID] = btnlabel

	return button
}
