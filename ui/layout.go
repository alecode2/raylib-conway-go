package ui

import (
	"fmt"
	"math"
)

func Size(root Element) {
	// Bottom-up phase: Fit sizing
	SizeAlongAxis(root)
	SizeAcrossAxis(root)

	// Top-down phase: Grow sizing
	ApplyGrowSizes(root, Horizontal)
	ApplyGrowSizes(root, Vertical)
}

func SizeAlongAxis(element Element) float32 {
	base := element.GetUIBase()
	axis := base.Direction
	sizing := getSizing(base, axis)
	fmt.Printf("Sizing for %s on axis %v: %v\n", base.ID, axis, sizing)

	// Fixed-size shortcut
	if sizing == SizingFixed {
		fmt.Printf("Fixed size for %s on axis %v: %f\n", base.ID, axis, getFixedSize(base, axis))
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	}

	// Leaf fallback
	if len(base.Children) == 0 && sizing != SizingGrow {
		fmt.Printf("No children, fixed size for %s on axis %v: %f\n", base.ID, axis, getFixedSize(base, axis))
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	}

	// Sum child sizes
	var total float32
	for _, child := range base.Children {
		childSize := SizeAlongAxis(child)
		fmt.Printf("Child size for %s: %f\n", child.GetUIBase().ID, childSize) // Debug line
		total += childSize
	}
	total += base.Gap * float32(max(0, len(base.Children)-1))
	total += getPadding(base, axis)

	// Clamp and assign
	clampedTotal := max(total, getMinSize(base, axis))
	setSize(base, axis, clampedTotal)
	fmt.Printf("Clamped size for %s: %f\n", base.ID, clampedTotal)
	return clampedTotal
}

func SizeAcrossAxis(element Element) float32 {
	base := element.GetUIBase()
	axis := getCrossAxis(base.Direction)
	sizing := getSizing(base, axis)
	fmt.Printf("Sizing for %s on axis %v: %v\n", base.ID, axis, sizing)

	if sizing == SizingFixed {
		fmt.Printf("Sizing for %s on axis %v: %v\n", base.ID, axis, sizing)
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	}

	if len(base.Children) == 0 && sizing != SizingGrow {
		fmt.Printf("No children, fixed size for %s on axis %v: %f\n", base.ID, axis, getFixedSize(base, axis))
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	}

	// Max child size
	var maxSize float32
	for _, child := range base.Children {
		size := SizeAcrossAxis(child)
		fmt.Printf("Child size on cross axis for %s: %f\n", child.GetUIBase().ID, size)
		if size > maxSize {
			maxSize = size
		}
	}
	maxSize += getPadding(base, axis)

	// Clamp and assign
	clampedMaxSize := max(maxSize, getMinSize(base, axis))
	setSize(base, axis, clampedMaxSize)
	fmt.Printf("Clamped size for %s on cross axis: %f\n", base.ID, clampedMaxSize)
	return clampedMaxSize
}

func ApplyGrowSizes(element Element, axis Axis) {
	base := element.GetUIBase()

	// Collect growable children
	var growable []Element
	usedSpace := float32(0)

	for _, child := range base.Children {
		childBase := child.GetUIBase()
		sizing := getSizing(childBase, axis)

		if sizing == SizingGrow {
			growable = append(growable, child)
		} else {
			usedSpace += getSize(childBase, axis)
		}
	}

	// Add up gaps and padding
	usedSpace += getPadding(base, axis)
	usedSpace += base.Gap * float32(max(0, len(base.Children)-1))

	remaining := getSize(base, axis) - usedSpace
	if remaining > 0 && len(growable) > 0 {
		GrowChildElements(base, growable, axis, remaining)
	}

	// Recurse top-down
	for _, child := range base.Children {
		ApplyGrowSizes(child, axis)
	}
}

func GrowChildElements(parent *UIBase, growable []Element, axis Axis, remaining float32) {
	for remaining > 0 {
		// Find smallest growable size
		smallest := getSize(growable[0].GetUIBase(), axis)
		secondSmallest := float32(math.Inf(1))
		widthToAdd := remaining

		for _, child := range growable {
			size := getSize(child.GetUIBase(), axis)

			if size < smallest {
				secondSmallest = smallest
				smallest = size
			} else if size > smallest {
				secondSmallest = min(secondSmallest, size)
				widthToAdd = secondSmallest - smallest
			}
		}

		// Determine step amount
		widthToAdd = min(widthToAdd, remaining/float32(len(growable)))

		for _, child := range growable {
			childBase := child.GetUIBase()
			if getSize(childBase, axis) == smallest {
				newSize := getSize(childBase, axis) + widthToAdd
				setSize(childBase, axis, newSize)
				remaining -= widthToAdd
			}
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

		fmt.Printf("Child %s position: X: %f, Y: %f\n", childBase.ID, childBase.Bounds.X, childBase.Bounds.Y)

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
