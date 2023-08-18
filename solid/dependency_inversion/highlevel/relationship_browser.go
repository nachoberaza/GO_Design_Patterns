package highlevel

import "github.com/nachoberaza/GO_Desing_Patterns/solid/dependency_inversion/lowlevel"

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*lowlevel.Person
}
