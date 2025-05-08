package ui

func Size(element Element) {
	base := element.GetUIBase()
	SizeAlongAxis(element, base.Direction)
	SizeAcrossAxis(element)
	for _, child := range base.Children {
		Size(child)
	}
}

func SizeAlongAxis(element Element, direction Axis) float32 {
	base := element.GetUIBase()

	var contentSize float32

	if (direction == Horizontal && base.WidthSizing == SizingFixed) ||
		(direction == Vertical && base.HeightSizing == SizingFixed) {
		// Use explicit size for fixed sizing
		if direction == Horizontal {
			contentSize = base.Width
		} else {
			contentSize = base.Height
		}
	} else if len(base.Children) == 0 {
		// Leaf nodes with fit/grow sizing
		if direction == Horizontal {
			contentSize = base.Width
		} else {
			contentSize = base.Height
		}
	} else {
		// Composite size from children + gaps
		for _, child := range base.Children {
			contentSize += SizeAlongAxis(child, direction)
		}
		contentSize += base.Gap * float32(max(0, len(base.Children)-1))
	}

	var total float32
	if direction == Horizontal {
		total = base.PaddingLeft + contentSize + base.PaddingRight
		if total < base.MinWidth {
			total = base.MinWidth
		}
	} else {
		total = base.PaddingTop + contentSize + base.PaddingBottom
		if total < base.MinHeight {
			total = base.MinHeight
		}
	}

	return total
}

func SizeAcrossAxis(element Element) {
	//TODO: Same as other function but with max size instead of sum of sizes
}

func Position(element Element) {
	for _, child := range element.GetUIBase().Children {
		Position(child)
	}
}

func Draw(element Element) {
	if d, ok := element.(Drawable); ok {
		d.Draw()
	}

	for _, child := range element.GetUIBase().Children {
		Draw(child)
	}
}
