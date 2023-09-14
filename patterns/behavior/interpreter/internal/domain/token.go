package domain

import "fmt"

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type TokenType int

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}
