package main

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/creational/builder/internal/directors"

func main() {
	director := directors.BuilderDirector{}
	director.DirectConstruction()
}
