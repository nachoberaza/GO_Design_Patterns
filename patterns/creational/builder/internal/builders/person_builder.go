package builders

import "github.com/baldurdevs/go_desing_patterns/patterns/creational/builder/internal/domain"

type PersonBuilder struct {
	person *domain.Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func (pb *PersonBuilder) Build() *domain.Person {
	return pb.person
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{personBuilder: pb}
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{personBuilder: pb}
}
