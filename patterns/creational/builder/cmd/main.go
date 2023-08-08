package main

import "github.com/baldurdevs/go_desing_patterns/patterns/creational/builder/internal/directors"

func main() {
	director := directors.BuilderDirector{}
	director.DirectConstruction()
}
