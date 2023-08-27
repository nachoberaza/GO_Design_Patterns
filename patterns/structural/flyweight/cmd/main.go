package main

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/cmd/shared"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/internal/factory"
)

func main() {
	game := domain.NewGame()

	//Add Terrorist
	game.AddTerrorist(shared.TerroristDressType)
	game.AddTerrorist(shared.TerroristDressType)
	game.AddTerrorist(shared.TerroristDressType)
	game.AddTerrorist(shared.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(shared.CounterTerroristDressType)
	game.AddCounterTerrorist(shared.CounterTerroristDressType)
	game.AddCounterTerrorist(shared.CounterTerroristDressType)

	dressFactoryInstance := factory.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
