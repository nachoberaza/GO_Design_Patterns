package abstractions

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/bridge/implementations"

type Computer interface {
	Print()
	SetPrinter(printer implementations.Printer)
}
