package domain

type Shape interface {
	GetType() string
	Accept(Visitor)
}
