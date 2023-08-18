package main

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/creational/prototype/internal/domain"
)

func main() {
	john := domain.Person{
		Name: "John",
		Address: &domain.Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
