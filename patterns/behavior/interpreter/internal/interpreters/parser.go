package interpreters

import (
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/interpreter/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/interpreter/internal/interfaces"
	"strconv"
)

func Parse(tokens []domain.Token) interfaces.Element {
	result := domain.BinaryOperation{}
	haveLhs := false
	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case domain.Int:
			n, _ := strconv.Atoi(token.Text)
			integer := domain.NewInteger(n)
			if !haveLhs {
				result.Left = integer
				haveLhs = true
			} else {
				result.Right = integer
			}
		case domain.Plus:
			result.Type = domain.Addition
		case domain.Minus:
			result.Type = domain.Subtraction
		case domain.Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == domain.Rparen {
					break
				}
			}
			// now j points to closing bracket, so
			// process subexpression without opening
			var subexp []domain.Token
			for k := i + 1; k < j; k++ {
				subexp = append(subexp, tokens[k])
			}
			element := Parse(subexp)
			if !haveLhs {
				result.Left = element
				haveLhs = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return &result
}
