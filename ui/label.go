package ui

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	UIBase
	Text      string
	Font      rl.Font
	FontSize  float32
	FontColor rl.Color
	TextAlign TextAlign
	Wrap      bool
	Spacing   float32
}

func NewLabel(text string, font rl.Font, fsize float32, fcolor rl.Color, tAlign TextAlign, wrap bool) *Label {
	return &Label{
		UIBase:    NewUIBase(),
		Text:      text,
		Font:      font,
		FontSize:  fsize,
		FontColor: fcolor,
		TextAlign: tAlign,
		Wrap:      wrap,
		Spacing:   float32(1), //Default spacing
	}
}

func (l *Label) GetUIBase() *UIBase {
	return &l.UIBase
}

func (l *Label) Draw() {
	font := l.Font
	fontSize := l.FontSize
	spacing := float32(1)

	// 1) Split into lines
	lines := []string{l.Text}
	if l.Wrap {
		lines = strings.Split(l.Text, "\n")
	}

	// 2) Measure each line
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
		infos[i] = lineInfo{
			text:   line,
			width:  sz.X,
			height: sz.Y,
		}
		if sz.X > maxWidth {
			maxWidth = sz.X
		}
		totalHeight += sz.Y + spacing
	}
	if len(lines) > 0 {
		totalHeight -= spacing // remove trailing spacing
	}

	// 3) Compute vertical start so the block is centered in the bounds
	yStart := l.Bounds.Y + (l.Bounds.Height-totalHeight)/2

	// 4) Draw each line with horizontal alignment
	for i, info := range infos {
		var xPos float32
		switch l.TextAlign {
		case AlignTextLeft:
			xPos = l.Bounds.X
		case AlignTextCenter:
			xPos = l.Bounds.X + (l.Bounds.Width-info.width)/2
		case AlignTextRight:
			xPos = l.Bounds.X + l.Bounds.Width - info.width
		}
		y := yStart + float32(i)*(info.height+spacing)
		rl.DrawTextEx(font, info.text, rl.NewVector2(xPos, y), fontSize, spacing, l.FontColor)
	}
}

func (l *Label) Measure(axis Axis) float32 {
	font := rl.GetFontDefault()
	size := l.FontSize
	txt := l.Text

	textSize := rl.MeasureTextEx(font, txt, size, 1)

	if axis == Horizontal {
		return textSize.X
	}
	return textSize.Y
}
