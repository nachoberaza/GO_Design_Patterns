package caller

import abstractCommand "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/command/internal/abstract_command"

type Button struct {
	Command abstractCommand.Command
}

func (b *Button) Press() {
	b.Command.Execute()
}
