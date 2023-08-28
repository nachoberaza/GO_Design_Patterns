package concrethandlers

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/handlers"
)

type Reception struct {
	next handlers.Department
}

func (r *Reception) Execute(p *domain.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next handlers.Department) {
	r.next = next
}
