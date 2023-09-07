package main

import (
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/command/internal/caller"
	concreteCommand "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/command/internal/concrete_command"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/command/internal/concrete_receiver"
)

func main() {
	tv := &concrete_receiver.Tv{}

	onCommand := &concreteCommand.OnCommand{
		Device: tv,
	}

	offCommand := &concreteCommand.OffCommand{
		Device: tv,
	}

	onButton := &caller.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &caller.Button{
		Command: offCommand,
	}
	offButton.Press()
}
