package cmp

import ui "conway/ui"

type Container struct {
	*ui.UIBase
}

func NewContainer() *Container {
	return &Container{
		UIBase: ui.NewUIBase(),
	}
}

func (c *Container) GetUIBase() *ui.UIBase {
	return c.UIBase
}

func (c *Container) Draw() {}
