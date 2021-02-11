package main

import (
	"fmt"

	"github.com/kmnkit/learngo/banking"
)

func main() {
	account := banking.NewAccount("Marco")
	account.Deposit(1000)
	fmt.Println(account.Balance())
	account.Withdraw(100)
	fmt.Println(account.Balance())

	err := account.Withdraw(901)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account.Balance())
	}
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("Ace")
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account.String())
}
