package domain

import (
	"container/list"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/observer/internal/observables"
)

// whenever a person catches a cold,
// a doctor must be called
type Person struct {
	observables.Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: observables.Observable{Subs: new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}
