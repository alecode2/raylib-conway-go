package ui

func Size(element Element) {
	base := element.GetUIBase()
	sizeAlong := SizeAlongAxis(element, base.Direction)
	sizeAcross := SizeAcrossAxis(element, base.Direction)

	if base.Direction == Horizontal {
		base.Bounds.Width = sizeAlong
		base.Bounds.Height = sizeAcross
	} else {
		base.Bounds.Width = sizeAcross
		base.Bounds.Height = sizeAlong
	}

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

func SizeAcrossAxis(element Element, direction Axis) float32 {
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
		// Maximum size from children
		contentSize = 0
		for _, child := range base.Children {
			childSize := SizeAcrossAxis(child, direction)

			if contentSize < childSize {
				contentSize = childSize
			}
		}
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

func Position(element Element) {
	base := element.GetUIBase()
	cursor := float32(0)

	//The axis in which we are stacking elements needs to have the current cursor kept in a variable
	if base.Direction == Horizontal {
		cursor = base.Bounds.X + base.PaddingLeft
	} else {
		cursor = base.Bounds.Y + base.PaddingTop
	}

	for i, child := range base.Children {
		childBase := child.GetUIBase()

		if base.Direction == Horizontal {
			childBase.Bounds.X = cursor
			childBase.Bounds.Y = base.Bounds.Y + base.PaddingTop
			cursor += childBase.Bounds.Width
		} else {
			childBase.Bounds.Y = cursor
			childBase.Bounds.X = base.Bounds.X + base.PaddingLeft
			cursor += childBase.Bounds.Height
		}
		//Only add gap if there is another child after the current one
		if i < len(base.Children)-1 {
			cursor += base.Gap
		}

		//Recurse down to the leafs
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
