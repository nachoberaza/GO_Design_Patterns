package lowlevel

type Person struct {
	Name string
}

type info struct {
	from     *Person
	Relation Relationship
	to       *Person
}
