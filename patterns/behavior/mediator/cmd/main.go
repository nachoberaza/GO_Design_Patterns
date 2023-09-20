package main

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/mediator/internal/domain"

func main() {
	room := domain.ChatRoom{}

	john := domain.NewPerson("John")
	jane := domain.NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi room")
	jane.Say("oh, hey john")

	simon := domain.NewPerson("Simon")
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage("Simon", "glad you could join us!")
}
