package main

import (
	"fmt"

	"github.com/kmnkit/learngo/banking"
)

func main() {
	account := banking.NewAccount("Marco")
	fmt.Println(account)
}
