package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Transactions interface {
// 	Deposit(amount int)
// 	Withdraw(amount int) error
// 	Balance() int
// }

// type Wallet struct {
// 	balance int
// }

// func main() {
// 	var wallet = Wallet{balance: 5}

// 	wallet.Deposit(10)
// 	wallet.Deposit(15)

// 	printBalance(&wallet)
// }

// func printBalance(wallet Transactions) {
// 	result := wallet.Balance()
// 	fmt.Println(result)
// }

// func (wallet *Wallet) Deposit(amount int) {
// 	wallet.balance += amount
// }

// func (wallet *Wallet) Balance() int {
// 	return wallet.balance
// }

// func (wallet *Wallet) Withdraw(amount int) error {
// 	if amount > wallet.balance {
// 		return errors.New("not enough money")
// 	}

// 	wallet.balance -= amount
// 	return nil
// }
