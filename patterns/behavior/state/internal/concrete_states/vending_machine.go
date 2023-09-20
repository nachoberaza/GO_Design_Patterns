package concrete_states

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/state/internal/state"
)

type VendingMachine struct {
	hasItem       state.State
	ItemRequested state.State
	hasMoney      state.State
	NoItem        state.State

	currentState state.State

	ItemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		ItemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		VendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		VendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		VendingMachine: v,
	}
	noItemState := &NoItemState{
		VendingMachine: v,
	}

	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.ItemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.NoItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) SetState(s state.State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.ItemCount = v.ItemCount + count
}
