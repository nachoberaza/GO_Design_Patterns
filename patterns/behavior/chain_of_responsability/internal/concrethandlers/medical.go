package concrethandlers

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/handlers"
)

type Medical struct {
	next handlers.Department
}

func (m *Medical) Execute(p *domain.Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next handlers.Department) {
	m.next = next
}
