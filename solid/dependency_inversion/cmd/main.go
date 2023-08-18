package main

import (
	"github.com/nachoberaza/GO_Desing_Patterns/solid/dependency_inversion/highlevel"
	"github.com/nachoberaza/GO_Desing_Patterns/solid/dependency_inversion/lowlevel"
)

func main() {
	parent := lowlevel.Person{Name: "John"}
	child1 := lowlevel.Person{Name: "Chris"}
	child2 := lowlevel.Person{Name: "Matt"}

	relationships := lowlevel.Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := highlevel.Research{Browser: &relationships}
	research.Investigate()
}
