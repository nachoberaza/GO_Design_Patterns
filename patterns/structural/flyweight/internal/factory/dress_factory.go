package factory

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/cmd/shared"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/internal/flyweightinterfaces"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/flyweight/internal/flyweightobjects"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]flyweightinterfaces.Dress),
	}
)

type DressFactory struct {
	DressMap map[string]flyweightinterfaces.Dress
}

func (d *DressFactory) GetDressByType(dressType string) (flyweightinterfaces.Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	switch dressType {
	case shared.TerroristDressType:
		d.DressMap[dressType] = flyweightobjects.NewTerroristDress()
		break
	case shared.CounterTerroristDressType:
		d.DressMap[dressType] = flyweightobjects.NewCounterTerroristDress()
		break
	default:
		return nil, fmt.Errorf("wrong dress type passed")
	}

	return d.DressMap[dressType], nil
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
