package cmp

import (
	"conway/event"
	ui "conway/ui"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputField struct {
	*ui.UIBase
	Texture    rl.Texture2D
	DrawConfig ui.DrawConfig
	Value      string
	Label      *Label
	OnSubmit   func(string)
	EventBus   *event.EventBus
}

func NewInputField(texture rl.Texture2D, style *ui.StyleSheet, label *Label, bus *event.EventBus) *InputField {
	base := ui.NewUIBase()
	base.Style = style

	input := &InputField{
		UIBase:     base,
		Texture:    texture,
		DrawConfig: ui.DrawConfig{Mode: ui.DrawModeSimple},
		Label:      label,
		EventBus:   bus,
	}

	ui.AddChild(input, label)

	ui.AddEventHandler(input, ui.EventRelease, func(e ui.UIEvent) {
		if input.EventBus != nil {
			input.EventBus.Emit(event.Event{Name: "focus_hex_input"})
		}
	})

	return input
}

func (in *InputField) GetUIBase() *ui.UIBase {
	return in.UIBase
}

func (in *InputField) SetText(s string) {
	in.Value = s
	in.Label.Text = s
}

func (in *InputField) AppendRune(r rune) {
	in.Value += string(r)
	in.Label.Text = in.Value
}

func (in *InputField) Backspace() {
	if len(in.Value) > 0 {
		in.Value = in.Value[:len(in.Value)-1]
		in.Label.Text = in.Value
	}
}

func (in *InputField) Draw() {
	if in.Texture.ID == 0 {
		fmt.Println("WARNING: InputField has no texture")
		return
	}

	tintVal := ui.ResolveStyle(in.UIBase, ui.Tint)
	tint, ok := tintVal.(rl.Color)
	if !ok {
		tint = rl.White
	}

	dest := in.Bounds
	switch in.DrawConfig.Mode {
	case ui.DrawModeSimple:
		src := rl.NewRectangle(0, 0, float32(in.Texture.Width), float32(in.Texture.Height))
		rl.DrawTexturePro(in.Texture, src, dest, rl.NewVector2(0, 0), 0, tint)

	case ui.DrawModeNineSlice:
		ui.DrawNineSlice(
			in.Texture,
			in.DrawConfig.NineSlice,
			dest,
			tint,
			in.DrawConfig.TileCenter,
			in.DrawConfig.TileEdges,
		)

	case ui.DrawModeTiled:
		src := rl.NewRectangle(0, 0, float32(in.Texture.Width), float32(in.Texture.Height))
		ui.TileTexture(in.Texture, src, dest, tint)

	default:
		fmt.Println("ERROR: Unknown DrawMode")
	}
}

func (in *InputField) IsHovered(mouse rl.Vector2) bool {
	return rl.CheckCollisionPointRec(mouse, in.Bounds)
}
