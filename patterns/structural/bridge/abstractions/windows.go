package abstractions

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/bridge/implementations"
)

type Windows struct {
	printer implementations.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p implementations.Printer) {
	w.printer = p
}
