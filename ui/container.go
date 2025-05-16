package ui

type Container struct {
	*UIBase
}

func NewContainer() *Container {
	return &Container{
		UIBase: NewUIBase(),
	}
}

func (c *Container) GetUIBase() *UIBase {
	return c.UIBase
}

func (c *Container) Draw() {
}
