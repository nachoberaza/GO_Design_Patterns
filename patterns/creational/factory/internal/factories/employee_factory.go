package factories

import "github.com/baldurdevs/go_desing_patterns/patterns/creational/factory/internal/domain"

type EmployeeFactory struct {
	position     string
	annualIncome int
}

func NewEmployeeFactory(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		position:     position,
		annualIncome: annualIncome,
	}
}

func (ef *EmployeeFactory) Create(name string) *domain.Employee {
	return &domain.Employee{Name: name, Position: ef.position, AnnualIncome: ef.annualIncome}
}
