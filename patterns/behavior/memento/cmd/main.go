package main

import (
	"fmt"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/memento/internal/domain"
)

func main() {
	ba := domain.NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 2:", ba)
	ba.Redo()
	fmt.Println("Redo:", ba)
}
