package main

import (
	"fmt"
	"log"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/facade/internal"
)

func main() {
	fmt.Println()
	walletFacade := internal.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
