package concrete_states

import "fmt"

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.VendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}
