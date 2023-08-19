package abstractions

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/bridge/implementations"
)

type Mac struct {
	printer implementations.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p implementations.Printer) {
	m.printer = p
}
