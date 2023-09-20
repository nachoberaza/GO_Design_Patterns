package observables

import (
	"container/list"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/observer/internal/observers"
)

type Observable struct {
	Subs *list.List
}

func (o *Observable) Subscribe(x observers.Observer) {
	o.Subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x observers.Observer) {
	for z := o.Subs.Front(); z != nil; z = z.Next() {
		if z.Value.(observers.Observer) == x {
			o.Subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.Subs.Front(); z != nil; z = z.Next() {
		z.Value.(observers.Observer).Notify(data)
	}
}
