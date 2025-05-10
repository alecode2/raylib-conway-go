package ui

import (
	//"fmt"
	"math"
)

func Size(root Element) {
	// Bottom-up phase: Fit sizing
	SizeRecursive(root)

	// Top-down phase: Grow sizing
	ApplyGrowSizes(root)
}

func SizeRecursive(element Element) {
	SizeAlongAxis(element)
	SizeAcrossAxis(element)

	for _, child := range element.GetUIBase().Children {
		SizeRecursive(child)
	}
}

func SizeAlongAxis(element Element) float32 {
	base := element.GetUIBase()
	axis := base.Direction
	sizing := getSizing(base, axis)

	// Fixed-size shortcut
	if sizing == SizingFixed {
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	}

	// Leaf fallback
	if len(base.Children) == 0 && sizing != SizingGrow {
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	}

	// Sum child sizes
	var total float32
	for _, child := range base.Children {
		childSize := SizeAlongAxis(child)
		total += childSize
	}
	total += base.Gap * float32(max(0, len(base.Children)-1))
	total += getPadding(base, axis)

	// Clamp and assign
	clampedTotal := max(total, getMinSize(base, axis))
	setSize(base, axis, clampedTotal)
	return clampedTotal
}

func SizeAcrossAxis(element Element) float32 {
	base := element.GetUIBase()
	axis := getCrossAxis(base.Direction)
	sizing := getSizing(base, axis)

	switch sizing {
	case SizingFixed:
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	case SizingGrow:
		size := getMinSize(base, axis)
		setSize(base, axis, size)
		return size
	default:
		// Compute max size from children
		var maxSize float32
		for _, child := range base.Children {
			size := SizeAcrossAxis(child)
			if size > maxSize {
				maxSize = size
			}
		}
		maxSize += getPadding(base, axis)

		clampedSize := max(maxSize, getMinSize(base, axis))
		setSize(base, axis, clampedSize)
		return clampedSize
	}
}

func ApplyGrowSizes(element Element) {
	base := element.GetUIBase()

	GrowAlongAxis(base, base.Direction)
	GrowAcrossAxis(base, base.Direction)

	for _, child := range base.Children {
		ApplyGrowSizes(child)
	}
}

func GrowAlongAxis(parent *UIBase, axis Axis) {
	// 1) Collect growable children and total used space
	var growable []Element
	used := float32(0)
	for _, child := range parent.Children {
		cb := child.GetUIBase()
		sizing := getSizing(cb, axis)
		if sizing == SizingGrow {
			growable = append(growable, child)
		}
		used += getSize(cb, axis)
	}

	// 2) Account for gaps
	gapTotal := parent.Gap * float32(max(0, len(parent.Children)-1))
	used += gapTotal

	// 3) Compute content space (excluding padding)
	total := getSize(parent, axis)
	padding := getPadding(parent, axis)
	contentSpace := total - padding
	remaining := contentSpace - used

	if remaining <= 0 || len(growable) == 0 {
		return
	}

	// 4) Distribute remaining
	for remaining > 0 {
		// find smallest
		smallest := getSize(growable[0].GetUIBase(), axis)
		second := float32(math.Inf(1))
		for _, child := range growable {
			sz := getSize(child.GetUIBase(), axis)
			if sz < smallest {
				second = smallest
				smallest = sz
			} else if sz > smallest {
				second = min(second, sz)
			}
		}

		// compute delta
		var delta float32
		if second == float32(math.Inf(1)) {
			delta = remaining / float32(len(growable))
		} else {
			delta = min(second-smallest, remaining/float32(len(growable)))
		}

		// apply
		for _, child := range growable {
			cb := child.GetUIBase()
			if getSize(cb, axis) == smallest {
				old := getSize(cb, axis)
				newSize := old + delta
				setSize(cb, axis, newSize)
				remaining -= delta
			}
		}
	}
}

func GrowAcrossAxis(parent *UIBase, axis Axis) {
	cross := getCrossAxis(axis)
	padding := getPadding(parent, cross)
	avail := getSize(parent, cross) - padding

	for _, child := range parent.Children {
		cb := child.GetUIBase()
		if getSizing(cb, cross) == SizingGrow {
			setSize(cb, cross, avail)
		}
	}
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

/*
UTILITY FUNCTIONS
*/
func clamp(val float32, min float32) float32 {
	if val < min {
		return min
	}
	return val
}

func getSizing(base *UIBase, axis Axis) SizingMode {
	if axis == Horizontal {
		return base.WidthSizing
	}
	return base.HeightSizing
}

func getMinSize(base *UIBase, axis Axis) float32 {
	if axis == Horizontal {
		return base.MinWidth
	}
	return base.MinHeight
}

func getFixedSize(base *UIBase, axis Axis) float32 {
	if axis == Horizontal {
		return base.Width
	}
	return base.Height
}

func getPadding(base *UIBase, axis Axis) float32 {
	if axis == Horizontal {
		return base.PaddingLeft + base.PaddingRight
	}
	return base.PaddingTop + base.PaddingBottom
}

func setSize(base *UIBase, axis Axis, value float32) {
	if axis == Horizontal {
		base.Width = value
		base.Bounds.Width = value
	} else {
		base.Height = value
		base.Bounds.Height = value
	}
}

func getCrossAxis(dir Axis) Axis {
	if dir == Horizontal {
		return Vertical
	}
	return Horizontal
}

func getSize(base *UIBase, axis Axis) float32 {
	if axis == Horizontal {
		return base.Width
	}
	return base.Height
}
