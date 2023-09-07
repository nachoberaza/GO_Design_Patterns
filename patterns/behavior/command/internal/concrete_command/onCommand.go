package concreteCommand

import (
	abstractReceiver "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/command/internal/abstract_receiver"
)

type OnCommand struct {
	Device abstractReceiver.Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
}
