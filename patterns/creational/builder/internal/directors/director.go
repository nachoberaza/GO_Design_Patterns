package directors

import (
	"fmt"

	"github.com/baldurdevs/go_desing_patterns/patterns/creational/builder/internal/builders"
)

type BuilderDirector struct {
	pb *builders.PersonBuilder
}

func (bd *BuilderDirector) DirectConstruction() {
	bd.pb = builders.NewPersonBuilder()

	personAddBuilder := bd.pb.Lives()
	personAddBuilder.At("123 London Road")
	personAddBuilder.In("London")
	personAddBuilder.WithPostCod("SW12BC")

	// Another way of object construction
	personJobBuilder := bd.pb.Works()
	personJobBuilder.At("Fabrikam").Earning(123000)

	person := bd.pb.Build()
	fmt.Println(*person)
}
