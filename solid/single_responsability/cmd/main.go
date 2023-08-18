package main

import (
	"github.com/nachoberaza/GO_Desing_Patterns/solid/single_responsability/journal"
)

func main() {
	journalObject := &journal.Journal{}
	journalObject.AddEntry("I cried today.")
	journalObject.AddEntry("I ate a bug")

	// Persistence
	p := &journal.Persistence{}
	err := p.SaveToFile(journalObject, "journal.txt")
	if err != nil {
		panic(err)
	}
}
