package Command

type OnCommand struct {
	Device Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
}
