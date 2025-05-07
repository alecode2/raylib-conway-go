package ui

type Axis int

const (
	Horizontal Axis = iota
	Vertical
)

func Size(element Element) {
	SizeAlongAxis(element)
	SizeAcrossAxis(element)
}

/*
func CalcSizeAlongAxis(el Element, axis Axis) float32 {
    if el is leaf:
        return el's intrinsic or fixed size on axis
    else:
        size := 0
        for child in el.GetChildren():
            size += CalcSizeAlongAxis(child, axis) + spacing/margin
        return size + el.padding/border
}
*/
func SizeAlongAxis(element Element) {
	if len(element.GetChildren()) == 0 {

	} else {

	}
}

/*
func CalcSizeAcrossAxis(el Element, axis Axis) float32 {
    if el is leaf:
        return el's intrinsic or fixed size across axis
    else:
        size := 0
        for child in el.GetChildren():
            size = max(size, CalcSizeAcrossAxis(child, axis))
        return size + el.padding/border
}
*/
func SizeAcrossAxis(element Element) {

}

/*
func LayoutPosition(el Element, origin rl.Vector2, axis Axis) {
    el.Bounds.X = origin.X
    el.Bounds.Y = origin.Y

    currentPos := origin
    for child in el.GetChildren():
        LayoutPosition(child, currentPos, axis)

        if axis == Horizontal:
            currentPos.X += child.Bounds.Width + spacing
        else:
            currentPos.Y += child.Bounds.Height + spacing
}
*/
func Position(element Element) {

}

func Draw(element Element) {
	if !element.IsVisible() {
		return
	}

	element.Draw()

	for _, child := range element.GetChildren() {
		Draw(child)
	}
}
