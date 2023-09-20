package main

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/visitor/internal/concrete_visitors"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/visitor/internal/domain"
)

func main() {
	square := &domain.Square{Side: 2}
	circle := &domain.Circle{Radius: 3}
	rectangle := &domain.Rectangle{
		L: 2,
		B: 3,
	}

	areaCalculator := &concrete_visitors.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &concrete_visitors.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
