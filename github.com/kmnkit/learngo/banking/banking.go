package banking

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
