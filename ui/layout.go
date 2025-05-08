package ui

type Axis int

const (
	Horizontal Axis = iota
	Vertical
)

func Size(element Element) {
	SizeAlongAxis(element)
	SizeAcrossAxis(element)
	for _, child := range element.GetUIBase().Children {
		Size(child)
	}
}

func SizeAlongAxis(element Element) {
	// TODO: Respect min/max constraints, auto-sizing, etc.
	// Use element.GetUIBase().Bounds.Width or Height depending on the axis
}

func SizeAcrossAxis(element Element) {
	// Similar to above but for the cross-axis
}

func Position(element Element) {
	// Calculate position based on layout rules
	for _, child := range element.GetUIBase().Children {
		Position(child)
	}
}

func Draw(element Element) {
	// Draw this element if it's a Drawable
	if d, ok := element.(Drawable); ok {
		d.Draw()
	}

	// Then draw its children
	for _, child := range element.GetUIBase().Children {
		Draw(child)
	}
}
