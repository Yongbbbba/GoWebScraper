package accounts

import "errors"

//struct는 python의 class와 object를 합친 것과 비슷

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// a Account는 receiver 라고 함
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance

}

var errNoMoney = errors.New("Can't wtihdraw, you are poor")

// Withdraw x amount from your account
// error handling
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner

}

// Owner of the acocunt
func (a Account) Owner() string {
	return a.owner

}
