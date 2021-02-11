package banking

import "errors"

var errNoMoney = errors.New("인출 할 수 없습니다. You are Poor!!!!")

// Account Struct
type Account struct {
	owner   string
	balance int
}

// NewAccount Create Account
func NewAccount(owner string) *Account {
	newAccount := Account{owner: owner, balance: 0}
	return &newAccount // 주소를 return
}

// Deposit 금액 추가
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw a amount of balance
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance of the Account
func (a Account) Balance() int {
	return a.balance
}
