package concrete_visitors

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/visitor/internal/domain"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) VisitForSquare(_ *domain.Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) VisitForCircle(_ *domain.Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) VisitForRectangle(_ *domain.Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
