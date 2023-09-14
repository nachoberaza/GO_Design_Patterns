package main

import (
	"fmt"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/interpreter/internal/interpreters"
)

func main() {
	input := "(13+4)-(12+1)"
	tokens := interpreters.Lex(input)
	fmt.Println(tokens)

	parsed := interpreters.Parse(tokens)
	fmt.Printf("%s = %d\n",
		input, parsed.Value())
}
