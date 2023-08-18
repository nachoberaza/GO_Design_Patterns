package main

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/solid/openclose/filter"
	"github.com/nachoberaza/GO_Desing_Patterns/solid/openclose/product"
)

func main() {
	apple := product.Product{
		Name:  "Apple",
		Color: product.Green,
		Size:  product.Small,
	}
	tree := product.Product{
		Name:  "Tree",
		Color: product.Green,
		Size:  product.Large,
	}
	house := product.Product{
		Name:  "House",
		Color: product.Blue,
		Size:  product.Large,
	}

	productsFilter := &filter.ProductFilter{
		ColorSpecification: &product.ColorSpecification{Color: product.Green},
		SizeSpecification:  &product.SizeSpecification{Size: product.Large},
	}

	products := []product.Product{apple, tree, house}
	fmt.Println("Green products:")
	for _, p := range productsFilter.FilterByColor(products) {
		fmt.Printf("* %s is green\n", p.Name)
	}

	fmt.Printf("\n")

	fmt.Println("Large products:")
	productsFilter.FilterBySize(products)
	for _, p := range productsFilter.FilterBySize(products) {
		fmt.Printf("* %s is large\n", p.Name)
	}

}
