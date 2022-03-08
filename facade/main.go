package main

import (
	"fmt"
	"log"
)

func main() {
	accountID := "abc"
	securityCode := 1234
	fmt.Println()
	walletFacade := newWalletFacade(accountID, securityCode)
	fmt.Println()
	err := walletFacade.addMoneyToWallet(accountID, securityCode, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet(accountID, securityCode, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
