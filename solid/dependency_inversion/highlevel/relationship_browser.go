package highlevel

import "github.com/baldurdevs/go_desing_patterns/solid/dependency_inversion/lowlevel"

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*lowlevel.Person
}
