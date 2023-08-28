package handlers

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/domain"

type Department interface {
	Execute(*domain.Patient)
	SetNext(Department)
}
