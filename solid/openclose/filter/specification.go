package filter

import "github.com/baldurdevs/go_desing_patterns/solid/openclose/product"

type Specification interface {
	IsSatisfied(p *product.Product) bool
}
