package main

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/decorator/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/decorator/internal/wrappers"
)

func main() {
	circle := domain.Circle{
		Radius: 30.0,
	}
	fmt.Println(circle.Render())

	// A wrapper
	redCircle := wrappers.ColoredShape{
		Shape: &domain.Circle{Radius: 2.0},
		Color: "Red",
	}
	fmt.Println(redCircle.Render())

	// Another wrapper into a wrapper
	transparentRedCircle := wrappers.TransparentShape{
		ColorShape: wrappers.ColoredShape{
			Shape: &domain.Circle{
				Radius: 30.0,
			},
			Color: "Red",
		},
		Transparency: 1.5,
	}
	fmt.Println(transparentRedCircle.Render())
}
