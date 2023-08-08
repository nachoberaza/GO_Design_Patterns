package lowlevel

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Relationships struct {
	Relations []info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range rs.Relations {
		if v.Relation == Parent &&
			v.from.Name == name {
			result = append(result, rs.Relations[i].to)
		}
	}

	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.Relations = append(rs.Relations, info{parent, Parent, child})
	rs.Relations = append(rs.Relations, info{child, Child, parent})
}
