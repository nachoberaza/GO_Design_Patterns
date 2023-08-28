package concrethandlers

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/handlers"
)

type Doctor struct {
	next handlers.Department
}

func (d *Doctor) Execute(p *domain.Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next handlers.Department) {
	d.next = next
}
