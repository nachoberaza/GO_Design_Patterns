package wrappers

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/decorator/internal/domain"
)

type ColoredShape struct {
	Shape domain.Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}
