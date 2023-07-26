package main

import (
	"github.com/baldurdevs/go_desing_patterns/solid_principles/single_responsability/journal"
	"sync"
)

func main() {
	journalObject := &journal.Journal{}
	journalObject.AddEntry("I cried today.")
	journalObject.AddEntry("I ate a bug")

	// Persistence
	p := &journal.Persistence{
		FileLocker: sync.Mutex{},
	}
	err := p.SaveToFile(journalObject, "./solid_principles/journal.txt")
	if err != nil {
		panic(err)
	}
}
