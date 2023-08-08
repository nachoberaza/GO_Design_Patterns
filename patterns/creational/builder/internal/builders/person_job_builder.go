package builders

/* Partial builder */
type PersonJobBuilder struct {
	personBuilder *PersonBuilder
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.personBuilder.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.personBuilder.person.AnnualIncome = annualIncome
	return pjb
}
