package cmp

import (
	ui "conway/ui"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	*ui.UIBase
	Text      string
	Font      rl.Font
	FontSize  float32
	TextAlign ui.TextAlign
	Wrap      bool
	Spacing   float32
}

func NewLabel(text string, font rl.Font, fontSize float32, style *ui.StyleSheet) *Label {
	base := ui.NewUIBase()

	return &Label{
		UIBase:    base,
		Text:      text,
		Font:      font,
		FontSize:  fontSize,
		TextAlign: ui.AlignTextLeft,
		Wrap:      false,
		Spacing:   1.0,
	}
}

func (l *Label) GetUIBase() *ui.UIBase {
	return l.UIBase
}

func (l *Label) Draw() {
	font := l.Font
	fontSize := l.FontSize
	spacing := l.Spacing

	// Resolve color from style
	colorVal := ui.ResolveStyle(l.UIBase, ui.Tint)
	fontColor, ok := colorVal.(rl.Color)
	if !ok {
		fontColor = rl.Black
	}

	// Split into lines (wrap currently just splits on \n)
	lines := []string{l.Text}
	if l.Wrap {
		lines = strings.Split(l.Text, "\n")
	}

	// Measure
	type lineInfo struct {
		text   string
		width  float32
		height float32
	}
	infos := make([]lineInfo, len(lines))
	maxWidth := float32(0)
	totalHeight := float32(0)

	for i, line := range lines {
		sz := rl.MeasureTextEx(font, line, fontSize, spacing)
		infos[i] = lineInfo{line, sz.X, sz.Y}
		if sz.X > maxWidth {
			maxWidth = sz.X
		}
		totalHeight += sz.Y + spacing
	}
	if len(lines) > 0 {
		totalHeight -= spacing
	}

	yStart := l.Bounds.Y + (l.Bounds.Height-totalHeight)/2

	for i, info := range infos {
		var xPos float32
		switch l.TextAlign {
		case ui.AlignTextLeft:
			xPos = l.Bounds.X
		case ui.AlignTextCenter:
			xPos = l.Bounds.X + (l.Bounds.Width-info.width)/2
		case ui.AlignTextRight:
			xPos = l.Bounds.X + l.Bounds.Width - info.width
		}
		y := yStart + float32(i)*(info.height+spacing)
		rl.DrawTextEx(font, info.text, rl.NewVector2(xPos, y), fontSize, spacing, fontColor)
	}
}

func (l *Label) Measure(axis ui.Axis) float32 {
	sz := rl.MeasureTextEx(l.Font, l.Text, l.FontSize, l.Spacing)
	if axis == ui.Horizontal {
		return sz.X
	}
	return sz.Y
}
