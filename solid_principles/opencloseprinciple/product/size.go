package product

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type SizeSpecification struct {
	Size Size
}

func (spec *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == spec.Size
}
