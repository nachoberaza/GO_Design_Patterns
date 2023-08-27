package domain

import (
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/internal/factory"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/internal/flyweightinterfaces"
)

type Player struct {
	dress      *flyweightinterfaces.Dress
	playerType string // it could be Terrorist o Counter-terrorist
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := factory.GetDressFactorySingleInstance().GetDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      &dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
