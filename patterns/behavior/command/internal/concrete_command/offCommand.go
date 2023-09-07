package concreteCommand

import abstractReceiver "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/command/internal/abstract_receiver"

type OffCommand struct {
	Device abstractReceiver.Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
}
