package iterators

import "github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/iterator/internal/domain"

type InOrderIterator struct {
	current       *domain.Node
	root          *domain.Node
	returnedStart bool
}

func NewInOrderIterator(root *domain.Node) Iterator {
	inOrderIterator := &InOrderIterator{
		current:       root,
		root:          root,
		returnedStart: false,
	}
	// move to the leftmost element
	for inOrderIterator.current.Left != nil {
		inOrderIterator.current = inOrderIterator.current.Left
	}
	return inOrderIterator
}

func (i *InOrderIterator) Reset() {
	i.current = i.root
	i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
	if i.current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true // can use first element
	}

	if i.current.Right != nil {
		i.current = i.current.Right
		for i.current.Left != nil {
			i.current = i.current.Left
		}
		return true
	} else {
		p := i.current.Parent
		for p != nil && i.current == p.Right {
			i.current = p
			p = p.Parent
		}
		i.current = p
		return i.current != nil
	}
}

func (it *InOrderIterator) GetCurrent() *domain.Node {
	return it.current
}
