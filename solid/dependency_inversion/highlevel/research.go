package highlevel

import "fmt"

type Research struct {
	Browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.Browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.Name)
	}
}
