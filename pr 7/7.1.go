// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
)

// банк
type BankAccount struct {
	accountNumber int
	holderName    string
	balance       float64
}

func Deposit(amount float64, Account *BankAccount) {
	Account.balance += amount
}

func Withdraw(amount float64, Account *BankAccount) error {
	if Account.balance < amount {
		return errors.New("недостаточно средств для снятия")
	} else {
		Account.balance -= amount
		return nil
	}
}

// пр нищита
func GetBalance(Account *BankAccount) float64 {
	return Account.balance
}

func main() {
	var Account BankAccount
	Deposit(100, &Account)
	Withdraw(50, &Account)
	fmt.Println("Ваш баланс:")
	fmt.Println(GetBalance(&Account))
}
