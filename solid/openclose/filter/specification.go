package filter

import "github.com/nachoberaza/GO_Desing_Patterns/solid/openclose/product"

type Specification interface {
	IsSatisfied(p *product.Product) bool
}
