package builders

/* Partial builder */
type PersonAddressBuilder struct {
	personBuilder *PersonBuilder
}

func (pab *PersonAddressBuilder) At(streetAddress string) {
	pab.personBuilder.person.StreetAddress = streetAddress
}

func (pab *PersonAddressBuilder) In(city string) {
	pab.personBuilder.person.City = city
}

func (pab *PersonAddressBuilder) WithPostCod(postCode string) {
	pab.personBuilder.person.PostCode = postCode
}
