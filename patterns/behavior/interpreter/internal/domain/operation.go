package domain

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/interpreter/internal/interfaces"

type Operation int

const (
	Addition Operation = iota
	Subtraction
)

type BinaryOperation struct {
	Type        Operation
	Left, Right interfaces.Element
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Subtraction:
		return b.Left.Value() + b.Right.Value()
	default:
		panic("Unsupported operation")
	}
}
