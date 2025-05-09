package ui

func Size(element Element) {
	base := element.GetUIBase()

	if base.Direction == Horizontal {
		base.Bounds.Width = SizeAlongAxis(element, Horizontal)
		base.Bounds.Height = SizeAcrossAxis(element, Vertical)
	} else {
		base.Bounds.Width = SizeAcrossAxis(element, Horizontal)
		base.Bounds.Height = SizeAlongAxis(element, Vertical)
	}

	for _, child := range base.Children {
		Size(child)
	}
}

func SizeAlongAxis(element Element, direction Axis) float32 {
	base := element.GetUIBase()
	// Fixed-size shortcut (still clamp at the end)
	if direction == Horizontal && base.WidthSizing == SizingFixed {
		return clamp(base.PaddingLeft+base.Width+base.PaddingRight, base.MinWidth)
	}
	if direction == Vertical && base.HeightSizing == SizingFixed {
		return clamp(base.PaddingTop+base.Height+base.PaddingBottom, base.MinHeight)
	}

	// Leaf nodes (no children)
	if len(base.Children) == 0 {
		if direction == Horizontal {
			return clamp(base.PaddingLeft+base.Width+base.PaddingRight, base.MinWidth)
		}
		return clamp(base.PaddingTop+base.Height+base.PaddingBottom, base.MinHeight)
	}

	// Step 1: Measure fixed and fit children
	fixedTotal := float32(0)
	growCount := 0
	childSizes := make([]float32, len(base.Children))

	for i, child := range base.Children {
		childBase := child.GetUIBase()

		// Pre-measure fit/fixed children
		sizing := childBase.WidthSizing
		if direction == Vertical {
			sizing = childBase.HeightSizing
		}

		if sizing == SizingGrow {
			growCount++
			continue
		}

		size := SizeAlongAxis(child, direction)
		childSizes[i] = size
		fixedTotal += size
	}

	// Step 2: Calculate remaining space
	totalGap := base.Gap * float32(max(0, len(base.Children)-1))
	padding := base.PaddingLeft + base.PaddingRight
	if direction == Vertical {
		padding = base.PaddingTop + base.PaddingBottom
	}

	containerSize := fixedTotal + totalGap + padding // without grow yet

	// Now determine container size based on sizing strategy
	if direction == Horizontal && base.WidthSizing == SizingFixed {
		containerSize = base.Width + padding
	}
	if direction == Vertical && base.HeightSizing == SizingFixed {
		containerSize = base.Height + padding
	}

	// Clamp to min
	minSize := base.MinWidth
	if direction == Vertical {
		minSize = base.MinHeight
	}
	if containerSize < minSize {
		containerSize = minSize
	}

	// Step 3: Distribute remaining space to grow children
	remaining := containerSize - fixedTotal - totalGap - padding
	growSize := float32(0)
	if growCount > 0 && remaining > 0 {
		growSize = remaining / float32(growCount)
	}

	// Step 4: Assign grow sizes and recurse
	for i, child := range base.Children {
		childBase := child.GetUIBase()

		sizing := childBase.WidthSizing
		if direction == Vertical {
			sizing = childBase.HeightSizing
		}

		if sizing == SizingGrow {
			if direction == Horizontal {
				childBase.Width = growSize
			} else {
				childBase.Height = growSize
			}
			// Recurse to propagate this down
			childSizes[i] = SizeAlongAxis(child, direction)
		}
	}

	// Step 5: Final total (include grow children now)
	total := float32(0)
	for _, sz := range childSizes {
		total += sz
	}
	total += totalGap + padding

	// Final clamp
	if direction == Horizontal && total < base.MinWidth {
		total = base.MinWidth
	}
	if direction == Vertical && total < base.MinHeight {
		total = base.MinHeight
	}
	return total
}

func SizeAcrossAxis(element Element, direction Axis) float32 {
	base := element.GetUIBase()

	// Early return for fixed size (still clamp at the end)
	if direction == Horizontal && base.WidthSizing == SizingFixed {
		return clamp(base.PaddingLeft+base.Width+base.PaddingRight, base.MinWidth)
	}
	if direction == Vertical && base.HeightSizing == SizingFixed {
		return clamp(base.PaddingTop+base.Height+base.PaddingBottom, base.MinHeight)
	}

	// Leaf node
	if len(base.Children) == 0 {
		if direction == Horizontal {
			return clamp(base.PaddingLeft+base.Width+base.PaddingRight, base.MinWidth)
		}
		return clamp(base.PaddingTop+base.Height+base.PaddingBottom, base.MinHeight)
	}

	// Step 1: Measure fixed/fit children
	fixedMax := float32(0)
	growCount := 0
	childSizes := make([]float32, len(base.Children))

	for i, child := range base.Children {
		childBase := child.GetUIBase()

		sizing := childBase.HeightSizing
		if direction == Horizontal {
			sizing = childBase.WidthSizing
		}

		if sizing == SizingGrow {
			growCount++
			continue
		}

		size := SizeAcrossAxis(child, direction)
		childSizes[i] = size

		if size > fixedMax {
			fixedMax = size
		}
	}

	// Step 2: Determine container size so far
	padding := base.PaddingTop + base.PaddingBottom
	minSize := base.MinHeight
	if direction == Horizontal {
		padding = base.PaddingLeft + base.PaddingRight
		minSize = base.MinWidth
	}
	containerSize := fixedMax + padding

	if containerSize < minSize {
		containerSize = minSize
	}

	// Step 3: Assign grow sizes
	// For across axis, grow children should match container size (minus padding)
	growSize := containerSize - padding
	if growSize < 0 {
		growSize = 0
	}

	for i, child := range base.Children {
		childBase := child.GetUIBase()

		sizing := childBase.HeightSizing
		if direction == Horizontal {
			sizing = childBase.WidthSizing
		}

		if sizing == SizingGrow {
			if direction == Horizontal {
				childBase.Width = growSize - childBase.PaddingLeft - childBase.PaddingRight
			} else {
				childBase.Height = growSize - childBase.PaddingTop - childBase.PaddingBottom
			}
			childSizes[i] = SizeAcrossAxis(child, direction)
		}
	}

	// Step 4: Final max size
	finalMax := float32(0)
	for _, sz := range childSizes {
		if sz > finalMax {
			finalMax = sz
		}
	}

	return clamp(finalMax+padding, minSize)
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

func clamp(val float32, min float32) float32 {
	if val < min {
		return min
	}
	return val
}
