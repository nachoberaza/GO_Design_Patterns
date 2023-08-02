package main

import (
	"fmt"

	"github.com/baldurdevs/go_desing_patterns/solid/liskov_substitution/shapes"
)

func main() {
	fmt.Printf("\n")

	rc := &shapes.Rectangle{
		Width:  2,
		Height: 3,
	}
	shapes.UseIt(rc)

	fmt.Printf("\n")

	sq := shapes.NewSquare(5)
	shapes.UseIt(sq)
}
