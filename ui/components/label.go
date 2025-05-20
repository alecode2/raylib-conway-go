package cmp

import (
	ui "conway/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strings"
)

type Label struct {
	*ui.UIBase
	Text      string
	Font      rl.Font
	FontSize  float32
	FontColor rl.Color
	TextAlign ui.TextAlign
	Wrap      bool
	Spacing   float32
}

func NewLabel(text string, font rl.Font, fsize float32, fcolor rl.Color, tAlign ui.TextAlign, wrap bool, spacing float32) *Label {
	return &Label{
		UIBase:    ui.NewUIBase(),
		Text:      text,
		Font:      font,
		FontSize:  fsize,
		FontColor: fcolor,
		TextAlign: tAlign,
		Wrap:      wrap,
		Spacing:   spacing,
	}
}

func (l *Label) GetUIBase() *ui.UIBase {
	return l.UIBase
}

func (l *Label) Draw() {
	font := l.Font
	fontSize := l.FontSize
	spacing := l.Spacing

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
		case ui.AlignTextLeft:
			xPos = l.Bounds.X
		case ui.AlignTextCenter:
			xPos = l.Bounds.X + (l.Bounds.Width-info.width)/2
		case ui.AlignTextRight:
			xPos = l.Bounds.X + l.Bounds.Width - info.width
		}
		y := yStart + float32(i)*(info.height+spacing)
		rl.DrawTextEx(font, info.text, rl.NewVector2(xPos, y), fontSize, spacing, l.FontColor)
	}
}

func (l *Label) Measure(axis ui.Axis) float32 {

	textSize := rl.MeasureTextEx(l.Font, l.Text, l.FontSize, l.Spacing)

	if axis == ui.Horizontal {
		return textSize.X
	}
	return textSize.Y
}
