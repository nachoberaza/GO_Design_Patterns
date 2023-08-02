package product

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type ColorSpecification struct {
	Color Color
}

func (spec *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == spec.Color
}
