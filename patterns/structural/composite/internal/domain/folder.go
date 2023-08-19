package domain

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/composite/internal/interfaces"
)

type Folder struct {
	Components []interfaces.Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c interfaces.Component) {
	f.Components = append(f.Components, c)
}
