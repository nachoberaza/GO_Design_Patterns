package concrethandlers

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/handlers"
)

type Cashier struct {
	next handlers.Department
}

func (c *Cashier) Execute(p *domain.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next handlers.Department) {
	c.next = next
}
