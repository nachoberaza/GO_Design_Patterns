package concrete_states

import (
	"fmt"
)

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}
func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.VendingMachine.ItemCount = i.VendingMachine.ItemCount - 1
	if i.VendingMachine.ItemCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.NoItem)
	} else {
		i.VendingMachine.SetState(i.VendingMachine.hasItem)
	}
	return nil
}
