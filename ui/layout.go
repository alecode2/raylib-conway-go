package ui

import (
	//"fmt"
	"fmt"
	"math"
	"strings"
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

	// Fixed‐size shortcut
	if sizing == SizingFixed {
		size := getFixedSize(base, axis)
		setSize(base, axis, size)
		return size
	}

	// Leaf fallback
	if len(base.Children) == 0 {

		size := getFixedSize(base, axis)
		if sizing == SizingGrow {
			size = getMinSize(base, axis)
		}
		setSize(base, axis, size)
		return size
	}

	// --- 1) Recurse into children first ---
	for _, child := range base.Children {
		SizeAlongAxis(child)
	}

	// --- 2) Now log and sum contributions ---
	fmt.Printf("[SizeAlongAxis Fit] ID=%s | axis=%v\n", base.ID, axis)
	var sumChildren float32
	for i, child := range base.Children {
		cb := child.GetUIBase()
		childSizing := getSizing(cb, axis)

		var childSize float32
		if childSizing == SizingGrow {
			childSize = getMinSize(cb, axis)
		} else {

			childSize = getSize(cb, axis)
		}

		fmt.Printf("  child[%d] %s mode=%v => sizeContribution=%.1f\n",
			i, cb.ID, childSizing, childSize)
		sumChildren += childSize
	}

	gapTotal := base.Gap * float32(max(0, len(base.Children)-1))
	padding := getPadding(base, axis)
	total := sumChildren + gapTotal + padding

	fmt.Printf(
		"  sumChildren=%.1f, gaps(total)=%.1f, padding=%.1f, total=%.1f\n",
		sumChildren, gapTotal, padding, total,
	)

	// Clamp and assign
	clamped := max(total, getMinSize(base, axis))
	setSize(base, axis, clamped)
	logSize("SizeAlongAxis return", base, axis, clamped)
	return clamped
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

	case SizingFit:
		fmt.Printf("[SizeAcrossAxis Fit] ID=%s | crossAxis=%v\n", base.ID, axis)
		var maxSize float32
		for _, child := range base.Children {
			cb := child.GetUIBase()
			childSizing := getSizing(cb, axis)

			var childSize float32
			if childSizing == SizingGrow {
				childSize = getMinSize(cb, axis)
			} else {
				childSize = getSize(cb, axis)
			}

			fmt.Printf("  child %s sizing=%v => sizeContribution=%.1f\n",
				cb.ID, childSizing, childSize)

			if childSize > maxSize {
				maxSize = childSize
			}
		}

		padding := getPadding(base, axis)
		total := maxSize + padding

		fmt.Printf("  maxChildBeforePadding=%.1f, padding=%.1f, total=%.1f\n",
			maxSize, padding, total)

		clampedSize := max(total, getMinSize(base, axis))
		setSize(base, axis, clampedSize)
		logSize("SizeAcrossAxis return", base, axis, clampedSize)
		return clampedSize
	}

	return 0
}

func ApplyGrowSizes(element Element) {
	base := element.GetUIBase()
	fmt.Printf("[ApplyGrowSizes] ID=%s\n", base.ID)

	GrowAlongAxis(base, base.Direction)
	GrowAcrossAxis(base, base.Direction)

	for _, child := range base.Children {
		ApplyGrowSizes(child)
	}
}

func GrowAlongAxis(parent *UIBase, axis Axis) {
	fmt.Printf("[GrowAlongAxis] Parent=%s | Axis=%v\n", parent.ID, axis)
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
	fmt.Printf("[GrowAcrossAxis] Parent=%s | CrossAxis=%v\n", parent.ID, cross)

	// If parent is Fit in the cross axis, do not grow children — it's up to them to determine their own size
	if getSizing(parent, cross) == SizingFit {
		fmt.Printf("[GrowAcrossAxis] SKIP due to SizingFit on %s\n", parent.ID)
		return
	}

	// Otherwise, children are allowed to grow to fill the parent's content area
	available := getSize(parent, cross) - getPadding(parent, cross)

	for _, child := range parent.Children {
		cb := child.GetUIBase()

		if getSizing(cb, cross) == SizingGrow {
			setSize(cb, cross, available)
		}
	}
}

func Position(element Element) {
	base := element.GetUIBase()

	// Build content box & child sizes
	main := base.Direction
	start, size := getContentBox(base, main)

	n := len(base.Children)
	sizes := make([]float32, n)
	var total float32
	for i, c := range base.Children {
		cb := c.GetUIBase()
		sizes[i] = getSize(cb, main)
		total += sizes[i]
	}

	// Distribute
	cursor, usedGap := distributeMain(
		start, size, n, total, base.Gap, base.MainAlign,
	)

	// Position each child
	for i, c := range base.Children {
		cb := c.GetUIBase()
		if main == Horizontal {
			cb.Bounds.X = cursor
		} else {
			cb.Bounds.Y = cursor
		}

		// Cross align
		alignCross(base, cb, base.CrossAlign)

		cursor += sizes[i] + usedGap
	}

	// Recurse
	for _, c := range base.Children {
		Position(c)
	}
}

func distributeMain(start, size float32, count int, totalChildren float32, gap float32, align MainAxisAlignment) (cursor, outGap float32) {
	free := size - totalChildren - gap*float32(max(0, count-1))
	switch align {
	case AlignStart:
		return start, gap
	case AlignEnd:
		return start + free, gap
	case AlignCenter:
		return start + free/2, gap
	case AlignSpaceBetween:
		if count > 1 {
			return start, free / float32(count-1)
		}
		return start + free/2, 0
	case AlignSpaceAround:
		if count > 0 {
			g := free / float32(count)
			return start + g/2, g
		}
		return start, 0
	}
	// fallback
	return start, gap
}

func alignCross(base *UIBase, child *UIBase, align CrossAxisAlignment) {
	axis := getCrossAxis(base.Direction)
	start, size := getContentBox(base, axis)
	childSize := getSize(child, axis)
	var pos float32
	switch align {
	case CrossAlignStart:
		pos = start
	case CrossAlignEnd:
		pos = start + (size - childSize)
	case CrossAlignCenter:
		pos = start + (size-childSize)/2
	}
	if axis == Horizontal {
		child.Bounds.X = pos
	} else {
		child.Bounds.Y = pos
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
	logSize("setSize", base, axis, value)
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

func getContentBox(base *UIBase, axis Axis) (start, size float32) {
	var padStart, padEnd float32
	if axis == Horizontal {
		start = base.Bounds.X
		padStart = base.PaddingLeft
		padEnd = base.PaddingRight
	} else {
		start = base.Bounds.Y
		padStart = base.PaddingTop
		padEnd = base.PaddingBottom
	}

	return start + padStart, getSize(base, axis) - (padStart + padEnd)
}

func logSize(prefix string, base *UIBase, axis Axis, size float32) {
	axisName := "Horizontal"
	if axis == Vertical {
		axisName = "Vertical"
	}
	fmt.Printf("[%s] ID=%s | Axis=%s | Size=%.2f\n", prefix, base.ID, axisName, size)
}

func PrintLayout(element Element, indent int) {
	base := element.GetUIBase()
	pad := strings.Repeat("  ", indent)
	fmt.Printf("%sID=%s W=%.1f H=%.1f\n", pad, base.ID, base.Bounds.Width, base.Bounds.Height)

	for _, child := range base.Children {
		PrintLayout(child, indent+1)
	}
}
