package filter

import "github.com/baldurdevs/go_desing_patterns/solid_principles/opencloseprinciple/product"

type ProductFilter struct {
	ColorSpecification *product.ColorSpecification
	SizeSpecification  *product.SizeSpecification
}

func (f *ProductFilter) FilterByColor(products []product.Product) []product.Product {
	result := make([]product.Product, 0)

	for i, p := range products {
		if f.ColorSpecification.IsSatisfied(&p) {
			result = append(result, products[i])
		}
	}

	return result
}

func (f *ProductFilter) FilterBySize(products []product.Product) []product.Product {
	result := make([]product.Product, 0)

	for i, p := range products {
		if f.SizeSpecification.IsSatisfied(&p) {
			result = append(result, products[i])
		}
	}

	return result
}
