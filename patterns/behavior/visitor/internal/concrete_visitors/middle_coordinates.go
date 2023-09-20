package concrete_visitors

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/visitor/internal/domain"
)

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) VisitForSquare(_ *domain.Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) VisitForCircle(_ *domain.Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) VisitForRectangle(_ *domain.Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
