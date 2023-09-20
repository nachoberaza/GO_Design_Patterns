package main

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/observer/internal/domain"

func main() {
	person := domain.NewPerson("Boris")
	doctor := domain.NewDoctor()

	person.Subscribe(doctor)

	person.CatchACold()
}
