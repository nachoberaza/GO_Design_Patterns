package filter

import "github.com/baldurdevs/go_desing_patterns/solid_principles/opencloseprinciple/product"

type Specification interface {
	IsSatisfied(p *product.Product) bool
}
