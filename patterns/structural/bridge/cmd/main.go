package main

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/bridge/abstractions"
	concreteimplementations "github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/bridge/concrete_implementations"
)

func main() {

	hpPrinter := &concreteimplementations.Hp{}
	epsonPrinter := &concreteimplementations.Epson{}

	macComputer := &abstractions.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &abstractions.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
