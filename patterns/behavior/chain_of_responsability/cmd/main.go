package main

import (
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/concrethandlers"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/chain_of_responsability/internal/domain"
)

func main() {

	cashier := &concrethandlers.Cashier{}

	//Set next for medical department
	medical := &concrethandlers.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &concrethandlers.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &concrethandlers.Reception{}
	reception.SetNext(doctor)

	patient := &domain.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
