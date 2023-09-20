package iterators

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/iterator/internal/domain"

type Iterator interface {
	Reset()
	MoveNext() bool
	GetCurrent() *domain.Node
}
