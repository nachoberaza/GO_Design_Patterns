package interpreters

import (
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/interpreter/internal/domain"
	"strings"
	"unicode"
)

func Lex(input string) []domain.Token {
	var result []domain.Token

	// not using range here
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, domain.Token{Type: domain.Plus, Text: "+"})
		case '-':
			result = append(result, domain.Token{Type: domain.Minus, Text: "-"})
		case '(':
			result = append(result, domain.Token{Type: domain.Lparen, Text: "("})
		case ')':
			result = append(result, domain.Token{Type: domain.Rparen, Text: ")"})
		default:
			sb := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
					i++
				} else {
					result = append(result, domain.Token{
						domain.Int, sb.String()})
					i--
					break
				}
			}
		}
	}
	return result
}
